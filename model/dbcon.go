package model

import (
	"fmt"
	"kakuninkun_server/logging"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB // インスタンス

// SQL接続とテーブル作成
func DBConnect() *gorm.DB {
	// .envから定数をプロセスの環境変数にロード
	err := godotenv.Load(".env") // エラーを格納
	if err != nil {              // エラーがあったら
		logging.ErrorLog("Error loading .env file", err)
		panic("Error loading .env file.")
	}
	// 環境変数から取得
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbDB := os.Getenv("MYSQL_DATABASE")

	// Mysqlに接続
	db, err = gorm.Open( // dbとエラーを取得
		"mysql", // dbの種類
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbDB), // 接続情報
	)
	if err != nil { // エラー処理
		fmt.Println("せつぞくできなかった")
		log.Fatal("Couldnt connect to the db server.", err)
	} else {
		fmt.Println("せつぞくできた")
		log.Println("Could connect to the db server.")
	}

	// テーブルがないなら自動で作成。 // gormがテーブル作成時にモデル名を複数形に、列名をスネークケースに。  // AutoMigrateは列情報の追加変更は反映するが列の削除は反映しない。
	db.AutoMigrate(
		&User{},
		&Kgroup{},
		&Company{},
	)

	// テスト用データ作成
	CreateKgroupTestData()
	CreateCompanyTestData()

	return db // 接続を返す
}

// 接続を取得
func GetDB() *gorm.DB {
	return db
}
