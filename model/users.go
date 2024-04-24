// モデルを構造体で定義  // gormがモデル名を複数形に、列名をスネークケースに。

package model

// ユーザテーブル
type User struct { // typeで型の定義, structは構造体
	ID             uint   `gorm:"primary_key;auto_increment;type:int(8)"` // 一意のid // json:"id"
	Name           string `gorm:"size:20"`                                // ユーザさんの名前
	DepartmentName string `gorm:"size:20"`                                // ？
	MailAddress    string `gorm:"size:64"`                                // メアド
	Password       string `gorm:"type:char(16)"`                          // パスワード
	Address        string `gorm:"size:100"`                               // 住所
	Situation      string `gorm:"size:5"`                                 // 状態
}

// 処理
