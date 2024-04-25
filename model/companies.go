// モデルを構造体で定義  // gormがモデル名を複数形に、列名をスネークケースに。

package model

// ユーザテーブル
type Company struct { // typeで型の定義, structは構造体
	CompanyNo   int    `gorm:"primary_key;AUTO_INCREMENT"` // 会社番号
	CompanyName string `gorm:"size:20;not null"`           // 会社名
}

// 処理
