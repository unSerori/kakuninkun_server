package model

// ユーザテーブル  // モデルを構造体で定義
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

// 処理
// 新規ユーザー作成
func CreateUser(newUser User) error {
	return db.Create(newUser).Error //　一行で。
	// if err := db.Create(newUser).Error; err != nil {
	// 	return err // 実行結果.Errorが存在してたら
	// }
	// return nil // エラーがない場合
}
