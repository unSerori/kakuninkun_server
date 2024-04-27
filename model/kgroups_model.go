package model

// 部署テーブル
type Kgroup struct { // typeで型の定義, structは構造体
	KgroupNo   int    `gorm:"primary_key;AUTO_INCREMENT"` // 部署番号
	KgroupName string `gorm:"size:20;not null"`           // 部署名
	CompanyNo  int    `gorm:"not null"`                   // 会社番号
}

// 処理
