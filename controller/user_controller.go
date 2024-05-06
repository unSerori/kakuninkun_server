package controller

import (
	"fmt"
	"kakuninkun_server/auth"
	"kakuninkun_server/logging"
	"kakuninkun_server/model"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

// register
func RegisterUser(c *gin.Context) {
	// JSONにマッピング
	// var reqBody map[string]interface{}             // リクエストを解析するためのmap
	// if err := c.ShouldBind(&reqBody); err != nil { // errがnilでないのでエラーハンドル
	// 	// エラーログ
	// 	logging.ErrorLog("Failed to mapping request JSON data.", err)
	// 	// レスポンス
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"srvResCode": 7001,                                // コード
	// 		"srvResMsg":  "Failed to mapping request JSON data.", // メッセージ
	// 		"srvResData": gin.H{},                             // データ
	// 	})
	// 	return // 早期リターンで終了
	// }
	// マップの中身を出力
	// fmt.Println("Map contents:")
	// for key, value := range reqBody {
	// 	fmt.Printf("%s: %v\n", key, value)
	// }

	// 構造体にマッピング
	var bUser model.User // 構造体のインスタンス
	if err := c.ShouldBindJSON(&bUser); err != nil {
		// エラーログ
		logging.ErrorLog("Failed to mapping request JSON data.", err)
		// レスポンス
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 7004,                                   // コード
			"srvResMsg":  "Failed to mapping request JSON data.", // メッセージ
			"srvResData": gin.H{},                                // データ
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
	// // マッピングされた構造体
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
	if token, err := auth.GenerateToken(bUser.Id); err != nil { // トークンを作成
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
	// 構造体にマッピング
	var bUser model.User // 構造体のインスタンス
	if err := c.ShouldBindJSON(&bUser); err != nil {
		// エラーログ
		logging.ErrorLog("Failed to mapping request JSON data.", err)
		// レスポンス
		c.JSON(http.StatusBadRequest, gin.H{
			"srvResCode": 7004,                                   // コード
			"srvResMsg":  "Failed to mapping request JSON data.", // メッセージ
			"srvResData": gin.H{},                                // データ
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

	tokenString, err := auth.GenerateToken(id)
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

// ユーザーの情報を取得
func UserProfile(c *gin.Context) {
	// ユーザーを特定する
	id, exists := c.Get("id")
	if !exists { // idがcに保存されていない。
		// エラーログ
		logging.ErrorLog("The id is not stored.", nil)
		// レスポンス
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 7013,                    // コード
			"srvResMsg":  "The id is not stored.", // メッセージ
			"srvResData": gin.H{},                 // データ
		})
		return
	}

	user, err := model.GetUserInfo(id.(int)) // ユーザーデータを取得
	if err != nil {                          // ユーザが見つからない。
		// エラーログ
		logging.ErrorLog("The condition specification may be correct, but the specified resource cannot be found.", err)
		// レスポンス
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 7003,                                                                                      // コード
			"srvResMsg":  "The condition specification may be correct, but the specified resource cannot be found.", // メッセージ
			"srvResData": gin.H{},                                                                                   // データ
		})
		return
	}

	// パスワードは返さない。  念のため。
	user.Password = ""
	if user.Password != "" { // 空文字になってなければ
		// エラーログ
		logging.ErrorLog("Failure to retrieve user data.", err)
		// レスポンス
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 7014,                             // コード
			"srvResMsg":  "Failure to retrieve user data.", // メッセージ
			"srvResData": gin.H{},                          // データ
		})
		return
	}

	// 取得に成功したのでユーザーデータを返す。
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 1003,                                          // コード
		"srvResMsg":  "Successful acquisition of user information.", // メッセージ
		"srvResData": gin.H{
			"userInfo": gin.H{
				"name":        user.Name,
				"id":          user.Id,
				"groupName":   user.GroupNo, // ここまで？
				"situation":   user.Situation,
				"mailAddress": user.MailAddress,
				"address":     user.Address,
				"company_no":  user.GroupNo,
			},
		}, // データ
	})
}

// ユーザー一覧の情報を取得
func UsersDataList(c *gin.Context) {
	// ユーザーとその所属会社を特定する
	id, exists := c.Get("id")
	if !exists { // idがcに保存されていない。
		// エラーログ
		logging.ErrorLog("The id is not stored.", nil)
		// レスポンス
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 7013,                    // コード
			"srvResMsg":  "The id is not stored.", // メッセージ
			"srvResData": gin.H{},                 // データ
		})
		return
	}

	fmt.Println("id: " + strconv.Itoa(id.(int)))

	compNo, err := model.GetCompanyNoById(id.(int)) // 会社番号を取得
	if err != nil {
		// エラーログ
		logging.ErrorLog("Failure to obtain company number.", nil)
		// レスポンス
		c.JSON(http.StatusInternalServerError, gin.H{
			"srvResCode": 7015,                                // コード
			"srvResMsg":  "Failure to obtain company number.", // メッセージ
			"srvResData": gin.H{},                             // データ
		})
		return
	}

	fmt.Println("compNo: " + strconv.Itoa(compNo))

	// 同じ会社のユーザー情報一覧を取得  // 1. 番号リストを取得して、それぞれで構造体データを取得。 // 2. 会社番号カラムが一致する行を取得
	users, err := model.GetUsersDataList(compNo) // type: []User
	if err != nil {
		fmt.Println("一覧取得に失敗")
		return
	}

	// スライスの各要素の構造体を、レスポンスに必要なフィールドだけ取得してjsonにする。

	adjustedUsers := []gin.H{}   // 短縮宣言  // var adjustedUsers []gin.H  // 宣言  // ginがレスポンスで使えるjson形式(:gin.H型)を要素とするスライス
	for _, user := range users { // 構造体スライスの要素に対してそれぞれ処理
		// 構造体から必要な分だけjson形式取り出す
		userJson := gin.H{ // json形式の要素を作成し、構造体から必要なフィールドを取得
			"name":      user.Name,
			"groupNo":   user.GroupNo,
			"situation": user.Situation,
		}
		// var userJson gin.H //userJson := make(gin.H)
		// userJson = make(gin.H)
		// userJson["name"] = user.Name
		// userJson["groupName"] = user.GroupNo
		// userJson["situation"] = user.Situation

		// スライスにぶちころがす
		adjustedUsers = append(adjustedUsers, userJson)
	}

	fmt.Println(adjustedUsers)

	// 取得出来たら返す
	c.JSON(http.StatusOK, gin.H{
		"srvResCode": 1002,                                                          // コード
		"srvResMsg":  "Successfully retrieved list of users matching the criteria.", // メッセージ
		"srvResData": gin.H{
			"userList": adjustedUsers,
		},
	})
}
