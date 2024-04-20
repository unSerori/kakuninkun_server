// package name
package main

// imports
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// TODO: モデルの定義
type Users struct { // typeで型の定義, structは構造体
	ID   uint   `gorm:"primary_key5"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// main method
func main() {
	// TODO: log func

	// エンジンを作成
	engine := gin.Default()

	// SQL接続
	// .envから定数をプロセスの環境変数にロード
	err := godotenv.Load(".env") // エラーを格納
	if err != nil {              // エラーがあったら
		panic("Error loding .env file")
	}
	// 環境変数から取得
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	//dbPort := os.Getenv("MYSQL_PORT")
	dbDB := os.Getenv("MYSQL_DATABASE")
	// 接続
	// Mysqlに接続
	db, err := gorm.Open( // dbとエラーを取得
		"mysql", // dbの種類
		fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbDB), // 接続情報
		// fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbDB), // 接続情報
	)
	if err != nil { // エラー処理
		fmt.Println("せつぞくできなかった")
		log.Fatal("Couldnt connect to the db server.", err)
	}
	defer db.Close() // defer文でこの関数が終了した際に破棄する

	db.AutoMigrate(&Users{})

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
