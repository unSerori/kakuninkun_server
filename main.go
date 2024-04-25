// package name
package main

// imports
import (
	"fmt"
	"kakuninkun_server/logging"
	"kakuninkun_server/model"
	"kakuninkun_server/route"
	"log"
)

// main method
func main() {
	// ログ設定を初期化
	logFile, err := logging.SetupLogging() // セットアップ
	if err != nil {                        // エラーチェック
		fmt.Printf("error opening file: %v\n", err)
	}
	defer logFile.Close() // 関数終了時に破棄

	log.Println("Start server!")

	// router設定されたengineを受け取る。
	router, err := route.GetRouter()
	if err != nil {
		fmt.Println(err) // エラー内容を出力し早期リターン ログ関連のエラーなのでログは出力しない
		return
	}

	// DB初期化
	model.DBConnect()

	// テンプレートと静的ファイルを読み込む
	router.LoadHTMLGlob("view/views/*.html")
	router.Static("/styles", "./view/views/styles") // クライアントがアクセスするURL, サーバ上のパス
	router.Static("/scripts", "./view/views/scripts")

	// 鯖起動
	router.Run(":4561")

	log.Println("Server successfully started.")
}
