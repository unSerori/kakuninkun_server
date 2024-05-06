package model

import (
	"errors"
)

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

// // ユーザーが存在するか確認
// func CheckUserExists(user User) error {
// 	return db.Where("mail_address = ?", user.MailAddress).First(&user).Error
// }

// パスワードが一致するか確認
func VerifyPass(user User) error {
	pass := user.Password // 入力されたパスワード
	var resultUser User   // 結果列を取得

	result := db.Where("mail_address = ?", user.MailAddress).Select("password").First(&resultUser) // mail_address列が入力メアドと一致する行のpassword列のみ取得しresultUserに取得
	if result.Error != nil {
		return result.Error
	}

	if resultUser.Password != pass {
		return errors.New("password mismatch") // 不一致
	}

	return nil // 一致
}

// メアドからidを取得
func GetIdByMail(user User) (int, error) {
	mail := user.MailAddress // 入力されたメアド
	var resultUser User      // 結果列を取得

	result := db.Where("mail_address = ?", mail).First(&resultUser) // メアドが一致する行を結果列として保存
	if result.Error != nil {
		return 0, result.Error
	}

	return resultUser.Id, nil // エラーなしの場合はidを返す。
}

// idが存在するか確かめる
func CfmId(id int) error {
	var user User
	return db.First(&user, "id = ?", id).Error // エラーなければnilが返る
}

// idからユーザー情報を取得
func GetUserInfo(id int) (*User, error) {
	var user User // 取得するユーザデータ
	if err := db.Select(
		"id, name, mail_address, address, situation, company_no, group_no",
	).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
	// 返り血の方とそれに伴った内部処理。
	// // idからユーザー情報を取得
	// func GetUserInfo(id int) (*User, error) {
	// 	var user User // 取得するユーザデータ
	// 	if err := db.Select(
	// 		"id, name, mail_address, address, situation, company_no, group_no",
	// 	).First(&user, "id = ?", id).Error; err != nil {
	// 		return nil, err
	// 	}
	// 	return &user, nil
	// }
	// // 返り血の型はポインタ変数
	// // 返り血の型はポインタ変数だが、結果をバインドする変数は構造体で初期化、あとでポインタのみ返す。
	// // First()の引数は必ずポインタ
	// // 返り血の型がポインタ変数なのでnilが返せる。
	// // 返り血の型がポインタ変数なので最初に作った構造体からポインタを返す。

	// // idからユーザー情報を取得
	//
	//	func GetUserInfo(id int) (User, error) {
	//		var user User // 取得するユーザデータ
	//		if err := db.Select(
	//			"id, name, mail_address, address, situation, company_no, group_no",
	//		).First(&user, "id = ?", id).Error; err != nil {
	//			return user, err
	//		}
	//		return user, nil
	//	}
	//
	// // 返り血の型は構造体のコピー
	// // 返り血の型は構造体のコピーなので結果をバインドする変数も構造体で初期化。
	// // First()の引数は必ずポインタ
	// // 返り血の型が構造体なのでnilが返せない。
	// // 返り血の型が構造体なのでバインドされた構造体をそのまま返す。
}

// idから簡易なユーザ情報を取得
func GetSimpleUserInfo(id int) (*User, error) {
	var user User // 取得するユーザデータ
	if err := db.Select(
		"id, name, mail_address, address, situation, company_no, group_no",
	).First(user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil

}
