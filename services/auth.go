package services

import (
	"fmt"
	"kakuninkun_server/logging"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// ユーザーidで認証トークンを生成
func GenerateToken(id int) (string, error) {
	// JWTのさまざまな情報を設定
	// .envから定数をプロセスの環境変数にロード
	err := godotenv.Load(".env") // エラーを格納
	if err != nil {              // エラーがあったら
		logging.ErrorLog("Error loading .env file", err)
		panic("Error loading .env file.")
	} // variable := os.Getenv("Key")
	// 環境変数から取得
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")                         // シークレットキー(署名鍵)を取得
	tokenLifeTime, err := strconv.Atoi(os.Getenv("JWT_TOKEN_LIFETIME")) // トークンの有効期限
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{ // クレーム部分
		"id":  id, // id。uuidにしたい。
		"exp": time.Now().Add(time.Second * time.Duration(tokenLifeTime)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)   // トークン(JWTを表す構造体)作成
	tokenString, err := token.SignedString([]byte(jwtSecretKey)) // []byte()でバイト型のスライスに変換し、SignedStringで署名。
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// トークン解析
func ParseToken(EnvVariables map[string]string, tokenString string) (*jwt.Token, error) {
	// 無名関数内で署名方法がHMACであるか確認し、正しければ秘密鍵を渡し、jwtトークンを解析する。
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logging.ErrorLog(fmt.Sprintf("Unexpected signature method: %v.", token.Header["alg"]), nil)
			return nil, fmt.Errorf("unexpected signature method: %v", token.Header["alg"])
		}
		return []byte(EnvVariables["JWT_SECRET_KEY"]), nil
	})
	if err != nil {
		return nil, err
	}

	// 正常に終われば解析されたトークンを渡す。
	return token, nil
}
