// package name
package main

// imports
import (
	"kakuninkun_server/model"
	"kakuninkun_server/route"
)

// main method
func main() {
	// router設定されたengineを受け取る。
	router := route.GetRouter()

	// DB初期化
	model.DBConnect()

	// テンプレートと静的ファイルを読み込む
	router.LoadHTMLGlob("view/views/*.html")
	router.Static("/styles", "./view/views/styles") // クライアントがアクセスするURL, サーバ上のパス
	router.Static("/scripts", "./view/views/scripts")

	// 鯖起動
	router.Run(":4561")
}
