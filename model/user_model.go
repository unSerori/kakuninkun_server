package model

import (
	"errors"
	"fmt"
	"strconv"
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
	fmt.Print("CfmId id: ")
	fmt.Println(id)
	var user User
	return db.First(&user, "id = ?", id).Error // エラーなければnilが返る
}

// idからユーザー情報を取得
func GetUserInfo(id int) (*User, error) {
	fmt.Println("AAAAAAAAA" + strconv.Itoa(id))
	var user User // 取得するユーザデータ

	// // 構造体の中身をチェック
	// st := reflect.TypeOf(user)  // 型を取得
	// sv := reflect.ValueOf(user) // 値を取得
	// // 構造体のフィールド数だけループ
	// for i := 0; i < st.NumField(); i++ {
	// 	fieldName := st.Field(i).Name                             // フィールド名を取得
	// 	fieldValue := sv.Field(i)                                 // フィールドの値を取得
	// 	fmt.Printf("%s: %v\n", fieldName, fieldValue.Interface()) // フィールド名と値を出力
	// }

	if err := db.Select(
		"id, name, mail_address, address, situation, company_no, group_no",
	).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// idから簡易なユーザ情報を取得
func GetSimpleUserInfo(id int) (*User, error) {
	var user *User // 取得するユーザデータ
	err := db.Select("id, name, mail_address, address, situation, company_no, group_no").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
