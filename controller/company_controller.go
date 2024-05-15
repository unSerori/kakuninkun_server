package controller

import (
	"kakuninkun_server/logging"
	"kakuninkun_server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 会社と企業一覧を整形して返す。
// 全体の流れ: 会社一覧をとってくる。それぞれの番号を取得して部署一覧から会社内の部署スライスを取得。
func CompList(c *gin.Context) {
	// 会社一覧を取得
	comps, err := model.CompList()
	if err != nil {
		// エラーログ
		logging.ErrorLog("Failure to retrieve company list.", nil)
		// レスポンス
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 7020,                                // コード
			"srvResMsg":  "Failure to retrieve company list.", // メッセージ
			"srvResData": gin.H{},                             // データ
		})
		return
	}

	adjustedComps := []gin.H{} // 調整後の会社構造体を生成

	// 1.それぞれの会社に名前などを設定する。2.それぞれの会社番号に対して、その番号を持つ(:その会社に所属する)部署一覧を当てはめる。
	for _, comp := range comps { // それぞれの会社に対して、
		// 会社の部署一覧をとる。
		groups, err := model.GetGroupListByComp(comp.CompanyNo)
		if err != nil { // compが持つ部署一覧が取れる
			// エラーログ
			logging.ErrorLog("Failure to obtain a list of group per company.", nil)
			// レスポンス
			c.JSON(http.StatusInternalServerError, gin.H{
				"srvResCode": 7021,                                             // コード
				"srvResMsg":  "Failure to obtain a list of group per company.", // メッセージ
				"srvResData": gin.H{},                                          // データ
			})
			return
		}

		// 軽量なJSONにマッピングしなおす
		simpleGroups := []gin.H{}      // JSONを要素にとるスライス
		for _, group := range groups { // 部署一覧のそれぞれを操作
			simpleGroups = append(simpleGroups,
				gin.H{ // いち部署。を新しいスライスの要素として追加
					"kgroupNo":   group.KgroupNo,
					"kgroupName": group.KgroupName,
				},
			)
		}

		// ひとつの会社データを完成させる
		adjustedComp := gin.H{
			"compNo":    comp.CompanyNo,   // これをもとに部署テーブルから検索した
			"compName":  comp.CompanyName, // 会社名
			"groupList": simpleGroups,     // 部署一覧
		}

		// 調整したそれぞれの会社情報をスライスにぶっこむ
		adjustedComps = append(adjustedComps, adjustedComp)
	}
	// 取得出来たら返す
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 1008,                                         // コード
		"srvResMsg":  "Successfully obtained a list of companies.", // メッセージ
		"srvResData": gin.H{
			"compList": adjustedComps,
		},
	})
	// 入れ子の構造体(:会社が部署を持つ形。jsonタグ設定済み)に入れてそのまま返してもいいかも
}
