// モデルを構造体で定義  // gormがモデル名を複数形に、列名をスネークケースに。

package model

// ユーザテーブル
type User struct { // typeで型の定義, structは構造体
	Id          int    `gorm:"primary_key;AUTO_INCREMENT;"` // 一意のid // json:"id"
	Name        string `gorm:"size:20;not null"`            // 名前
	MailAddress string `gorm:"size:64;not null;unique"`     // メアド
	Password    string `gorm:"size:16;not null;unique"`     // パスワード
	Address     string `gorm:"size:100;not null"`           // 住所
	Situation   string `gorm:"size:5"`                      // 状況
	CompanyNo   int    `gorm:"not null"`                    // 会社番号
	GroupNo     int    `gorm:"type:int(10);not null"`       // 部署番号
}

// 会社テーブル
type Company struct { // typeで型の定義, structは構造体
	CompanyNo   int    `gorm:"primary_key;AUTO_INCREMENT"` // 会社番号
	CompanyName string `gorm:"size:20;not null"`           // 会社名
}

// 部署テーブル
type Kgroup struct { // typeで型の定義, structは構造体
	KgroupNo   int    `gorm:"primary_key;AUTO_INCREMENT"` // 部署番号
	KgroupName string `gorm:"size:20;not null"`           // 部署名
	CompanyNo  int    `gorm:"not null"`                   // 会社番号
}
