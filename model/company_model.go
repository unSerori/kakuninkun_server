package model

// 会社テーブル
type Company struct { // typeで型の定義, structは構造体
	CompanyNo   int    `gorm:"primary_key;AUTO_INCREMENT"` // 会社番号
	CompanyName string `gorm:"size:20;not null"`           // 会社名
}

// 処理
// テストデータ作成
func CreateCompanyTestData() {
	kg := &Company{
		1,
		"AComp",
	}
	db.Create(kg)
	kg = &Company{
		2,
		"BComp",
	}
	db.Create(kg)
}
