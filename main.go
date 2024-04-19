// package name
package main

// imports
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// main method
func main() {
	// エンジンを作成
	engine := gin.Default()

	// endpoints
	// test
	engine.GET("/", func(c *gin.Context) { // GETメソッド("/route_path", ハンドラ関数(引数にリクエストとレスポンスに関する情報や操作を行うためのインタフェースであるgin.Context型のオブジェクトを受け取る))
		c.JSON(http.StatusOK, gin.H{ // bodyがJSON形式のレスポンスを返す
			"message": "hello go server!",
		})
	})

	// 起動
	engine.Run(":4561")
}
