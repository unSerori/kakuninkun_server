package model

// 部署テーブル
type Kgroup struct { // typeで型の定義, structは構造体
	KgroupNo   int    `gorm:"primary_key;AUTO_INCREMENT"` // 部署番号
	KgroupName string `gorm:"size:20;not null"`           // 部署名
	CompanyNo  int    `gorm:"not null"`                   // 会社番号
}

// 処理
// テストデータ作成
func CreateKgroupTestData() {
	kg1 := &Kgroup{
		1,
		"人事部",
		1,
	}
	db.Create(kg1)
	kg2 := &Kgroup{
		2,
		"人事部",
		2,
	}
	db.Create(kg2)
	kg3 := &Kgroup{
		KgroupNo:   3, // プライマリーキーを指定しないと自動で作成
		KgroupName: "情報技術部",
		CompanyNo:  1,
	}
	db.Create(kg3)
}
