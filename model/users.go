// モデルを構造体で定義  // gormがモデル名を複数形に、列名をスネークケースに。

package model

// ユーザテーブル
type User struct { // typeで型の定義, structは構造体
	Id          int    `gorm:"primary_key;AUTO_INCREMENT;"` // 一意のid // json:"id"
	Name        string `gorm:"size:20;not null"`            // 名前
	GroupNo     int    `gorm:"type:int(10)"`                //
	MailAddress string `gorm:"size:64;not null"`            // メアド
	Password    string `gorm:"size:16;not null"`            // パスワード
	Address     string `gorm:"size:100;not null"`           // 住所
	Situation   string `gorm:"size:5"`                      // 状況
}

// 処理
