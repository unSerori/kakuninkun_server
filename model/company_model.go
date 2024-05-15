package model

import (
	"fmt"
	"reflect"
)

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

// 会社一覧を取得する
func CompList() ([]Company, error) {
	var comps []Company // 取得したデータをマッピングする構造体 複数あるのでスライス
	if err := db.Find(&comps).Error; err != nil {
		return nil, err
	}
	for _, comp := range comps {
		fmt.Println("=======================")
		// 構造体の中身をチェック
		st := reflect.TypeOf(comp)  // 型を取得
		sv := reflect.ValueOf(comp) // 値を取得
		// 構造体のフィールド数だけループ
		for i := 0; i < st.NumField(); i++ {
			fieldName := st.Field(i).Name                             // フィールド名を取得
			fieldValue := sv.Field(i)                                 // フィールドの値を取得
			fmt.Printf("%s: %v\n", fieldName, fieldValue.Interface()) // フィールド名と値を出力
		}
	}

	return comps, nil // スライスを返す。
}

// 指定された会社番号から会社名を返す
func GetCompName(companyNo int) (string, error) {
	var comp Company // 取得したデータをマッピングする構造体
	if err := db.Find(&comp, companyNo).Error; err != nil {
		return "", err
	}
	return comp.CompanyName, nil
}
