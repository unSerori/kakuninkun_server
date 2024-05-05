package controller

import (
	"fmt"
	"kakuninkun_server/logging"
	"kakuninkun_server/model"
	"kakuninkun_server/services"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

// register
func RegisterUser(c *gin.Context) {
	// JSONにバインド
	// var reqBody map[string]interface{}             // リクエストを解析するためのmap
	// if err := c.ShouldBind(&reqBody); err != nil { // errがnilでないのでエラーハンドル
	// 	// エラーログ
	// 	logging.ErrorLog("Failed to bind request JSON data.", err)
	// 	// レスポンス
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"srvResCode": 7001,                                // コード
	// 		"srvResMsg":  "Failed to bind request JSON data.", // メッセージ
	// 		"srvResData": gin.H{},                             // データ
	// 	})
	// 	return // 早期リターンで終了
	// }
	// マップの中身を出力
	// fmt.Println("Map contents:")
	// for key, value := range reqBody {
	// 	fmt.Printf("%s: %v\n", key, value)
	// }

	// 構造体にバインド
	var bUser model.User // 構造体のインスタンス
	if err := c.ShouldBindJSON(&bUser); err != nil {
		// エラーログ
		logging.ErrorLog("Failed to bind request JSON data.", err)
		// レスポンス
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 7004,                                // コード
			"srvResMsg":  "Failed to bind request JSON data.", // メッセージ
			"srvResData": gin.H{},                             // データ
		})
		return // 早期リターンで終了
	}
	// // 構造体の中身をチェック
	// st := reflect.TypeOf(bUser)  // 型を取得
	// sv := reflect.ValueOf(bUser) // 値を取得
	// // 構造体のフィールド数だけループ
	// for i := 0; i < st.NumField(); i++ {
	// 	fieldName := st.Field(i).Name                             // フィールド名を取得
	// 	fieldValue := sv.Field(i)                                 // フィールドの値を取得
	// 	fmt.Printf("%s: %v\n", fieldName, fieldValue.Interface()) // フィールド名と値を出力
	// }

	// 登録用関数に渡す構造体を作成
	bUser.Situation = "安否確認中"

	// GolangのJSONエンコードデコードさんぷる
	// // 構造体->JSONバイト列->mapに変換してみる
	// fmt.Println("++++++++++++")
	// // バインドされた構造体
	// //var bUser model.User
	// fmt.Println(bUser.Password) // 構造体のフィールドにアクセス
	// // 構造体をJSONバイト列(:ASCII)にエンコード
	// mbUser, _ := json.Marshal(bUser) // 変換
	// fmt.Println(mbUser)              // バイト列として表現
	// fmt.Println(string(mbUser))      // 文字列として表現
	// // バイト列をMapにデコード
	// var umbUser map[string]interface{}      // mapの宣言
	// err := json.Unmarshal(mbUser, &umbUser) // 変換
	// _ = err                                 // エラー処理
	// fmt.Println(umbUser["Password"])        // mapの値をキーから参照する
	// fmt.Println("------------")
	// // map->JSONバイト列->構造体に戻してみる
	// fmt.Println("++++++++++++")
	// // 変換されたMap
	// fmt.Println(umbUser["Password"]) // mapの値をキーから参照する
	// // mapをJSONバイト列にエンコード
	// mumbUser, _ := json.Marshal(umbUser) // 変換
	// fmt.Println(mumbUser)                // バイト列として表現
	// fmt.Println(string(mumbUser))        // 文字列として表現
	// // JSONバイト列を構造体にデコード
	// var smumbUser model.User                   // インスタンス作成
	// err = json.Unmarshal(mumbUser, &smumbUser) // 変換
	// _ = err                                    // エラー処理
	// fmt.Println(bUser.Password)                // 構造体のフィールドにアクセス
	// fmt.Println("------------")

	// 構造体をレコード作成処理に渡す
	if err := model.CreateUser(bUser); err != nil {
		// エラー処理
		switch err.(*mysql.MySQLError).Number {
		case 1062: // 一意性制約違反
			// エラーログ
			logging.ErrorLog("There is already a user with the same primary key. Uniqueness constraint violation.", err)
			// レスポンス
			c.JSON(http.StatusInternalServerError, gin.H{
				"srvResCode": 7007,                                                                                  // コード
				"srvResMsg":  "There is already a user with the same primary key. Uniqueness constraint violation.", // メッセージ
				"srvResData": gin.H{},                                                                               // データ
			})
		default:
			// エラーログ
			logging.ErrorLog("Some problems with db registration of new users.", err)
			// レスポンス
			c.JSON(http.StatusInternalServerError, gin.H{
				"srvResCode": 7006,                                               // コード
				"srvResMsg":  "Some problems with db registration of new users.", // メッセージ
				"srvResData": gin.H{},                                            // データ
			})
		}
		return // 終了
	}

	//　成功
	if token, err := services.GenerateToken(bUser.Id); err != nil { // トークンを作成
		// エラーログ
		logging.ErrorLog("Failed to generate authentication token.", err)
		// レスポンス
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 7010,                                      // コード
			"srvResMsg":  "Failed to generate authentication token", // メッセージ
			"srvResData": gin.H{},                                   // データ
		})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"srvResCode": 1004,                            // コード
			"srvResMsg":  "Successful user registration.", // メッセージ
			"srvResData": gin.H{
				"authenticationToken": token,
			}, // データ
		})
	}

}

// login
func Login(c *gin.Context) {
	/*
			    {
		      "mailAddress": "hogeta@gmail.com",
		      "password": "C@h"
		    }
	*/
	// リクエストからログイン情報を取得
	// 構造体にバインド
	var bUser model.User // 構造体のインスタンス
	if err := c.ShouldBindJSON(&bUser); err != nil {
		// エラーログ
		logging.ErrorLog("Failed to bind request JSON data.", err)
		// レスポンス
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 7004,                                // コード
			"srvResMsg":  "Failed to bind request JSON data.", // メッセージ
			"srvResData": gin.H{},                             // データ
		})
		return // 早期リターンで終了
	}
	// 構造体の中身をチェック
	st := reflect.TypeOf(bUser)  // 型を取得
	sv := reflect.ValueOf(bUser) // 値を取得
	// 構造体のフィールド数だけループ
	for i := 0; i < st.NumField(); i++ {
		fieldName := st.Field(i).Name                             // フィールド名を取得
		fieldValue := sv.Field(i)                                 // フィールドの値を取得
		fmt.Printf("%s: %v\n", fieldName, fieldValue.Interface()) // フィールド名と値を出力
	}

	// // ユーザーが存在するか確認
	// if err := model.CheckUserExists(bUser); err != nil {
	// 	// エラーログ
	// 	logging.ErrorLog("User not found.", err)
	// 	// レスポンス
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"srvResCode": 7009,              // コード
	// 		"srvResMsg":  "User not found.", // メッセージ
	// 		"srvResData": gin.H{},           // データ
	// 	})
	// 	return
	// }
	// パスワードが一致するか確認
	if err := model.VerifyPass(bUser); err != nil {
		// エラーログ
		logging.ErrorLog("Password does not match.", err)
		// レスポンス
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 7010,                       // コード
			"srvResMsg":  "Password does not match.", // メッセージ
			"srvResData": gin.H{},                    // データ
		})
		return
	}

	// メールアドレスから検索したユーザーidをもとにトークンを作成
	id, err := model.GetIdByMail(bUser)
	if err != nil {
		// エラーログ
		logging.ErrorLog("Failure to obtain user ID.", err)
		// レスポンス
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 7010,                         // コード
			"srvResMsg":  "Failure to obtain user ID.", // メッセージ
			"srvResData": gin.H{},                      // データ
		})
		return
	}

	tokenString, err := services.GenerateToken(id)
	if err != nil {
		// エラーログ
		logging.ErrorLog("Failed to generate authentication token.", err)
		// レスポンス
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 7010,                                      // コード
			"srvResMsg":  "Failed to generate authentication token", // メッセージ
			"srvResData": gin.H{},                                   // データ
		})
		return
	}

	// パスワードが一致した場合
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 1005,                // コード
		"srvResMsg":  "Successful login.", // メッセージ
		"srvResData": gin.H{
			"authenticationToken": tokenString,
		}, // データ
	})
}

func UserProfile(c *gin.Context) {
	// ユーザーを特定する

}
