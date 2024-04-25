// モデルを構造体で定義  // gormがモデル名を複数形に、列名をスネークケースに。

package model

// ユーザテーブル
type Kgroup struct { // typeで型の定義, structは構造体
	KgroupNo   int    `gorm:"primary_key;AUTO_INCREMENT"` // グループ番号
	KgroupName string `gorm:"size:20;not null"`           // グループ名
}

// 処理
