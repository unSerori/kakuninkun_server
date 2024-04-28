package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// POSTリクエストボディをバインドするための構造体
// /register
type BoundRegister struct {
	CompanyName string `json:"companyName"`
	Name        string `json:"name"`
	Id          int    `json:"id"`
	MailAddress string `json:"mailAddress"`
	Address     string `json:"address"`
	Password    string `json:"password"`
}

// register
func RegisterUser(c *gin.Context) {
	var bRegister BoundRegister                          // 構造体のインスタンス
	if err := c.ShouldBindJSON(&bRegister); err != nil { // errがnilでないのでエラーハンドル
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 7001,                                // コード
			"srvResMsg":  "Failed to bind request JSON data.", // メッセージ
			"srvResErr":  err.Error(),                         // エラー内容
			"srvResData": gin.H{},                             // データ
		})
		return // 終了
	}

	// バインド後の処理
	println(bRegister.Password)

	// 構造体->JSONバイト列->mapに変換してみる
	mRegister, _ := json.Marshal(bRegister)       // 構造体をJSONバイト列(:ASCII)にエンコード
	fmt.Println(string(mRegister))                // 文字列に変換
	var umRegister map[string]interface{}         // mapを宣言
	err := json.Unmarshal(mRegister, &umRegister) // mapに変換
	_ = err
	fmt.Println(umRegister["password"]) // mapのキー指定で参照

	// map->JSONバイト列->構造体に戻してみる
	remRegister, _ := json.Marshal(umRegister) // mapをJSONバイト列に。
	fmt.Println(string(remRegister))           // 文字列に変換
	var jbRegister BoundRegister               // 構造体インスタンス
	json.Unmarshal(remRegister, &jbRegister)   // JSONバイト列を構造体にデコード
	fmt.Println(jbRegister.Address)
	// .Marshal()によるmap->構造体の直接変換はできない。

	// 構造体をレコード作成処理に渡す
	// record := model.User{
	// 	TODO: しょきか
	// }

	c.JSON(http.StatusCreated, gin.H{
		"srvResCode": 1004,                            // コード
		"srvResMsg":  "Successful user registration.", // メッセージ
		"srvResData": gin.H{},                         // データ
	})
}
