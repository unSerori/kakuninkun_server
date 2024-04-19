// package name
package main

// imports
import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// main method
func main() {
	// エンジンを作成
	engine := gin.Default()

	// テンプレートと静的ファイルを読み込む
	engine.LoadHTMLGlob("views/*.html")
	engine.Static("/styles", "./views/styles") // クライアントがアクセスするURL, サーバ上のパス
	engine.Static("/scripts", "./views/scripts")

	// endpoints
	// test
	engine.GET("/", func(c *gin.Context) { // GETメソッド("/route_path", ハンドラ関数(引数にリクエストとレスポンスに関する情報や操作を行うためのインタフェースであるgin.Context型のオブジェクトを受け取る))
		c.HTML(http.StatusOK, "index.html", gin.H{
			"topTitle":  "Route /",                                                            // ヘッダ内容 h1
			"mainTitle": "Hello.",                                                             // メインのタイトル h2
			"time":      time.Now(),                                                           // 時刻
			"message":   "This is an API server written in Golang for safety check purposes.", // message
		})
		// c.JSON(http.StatusOK, gin.H{  // bodyがJSON形式のレスポンスを返す
		// 	"message": "hello go server!",
		// })
	})

	// 起動
	engine.Run(":4561")
}
