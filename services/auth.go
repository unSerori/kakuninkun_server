package services

import (
	"errors"
	"fmt"
	"kakuninkun_server/logging"
	"kakuninkun_server/model"
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

	// クレーム部分
	claims := jwt.MapClaims{
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

// トークン解析検証
func ParseToken(tokenString string) (*jwt.Token, error) {
	// .envから定数をプロセスの環境変数にロード
	err := godotenv.Load(".env") // エラーを格納
	if err != nil {              // エラーがあったら
		logging.ErrorLog("Error loading .env file", err)
		panic("Error loading .env file.")
	}

	// 署名が正しければ、解析用の鍵を使う。(無名関数内で署名方法がHMACであるか確認し、HMACであれば秘密鍵を渡し、jwtトークンを解析する。)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // 署名を確認
			logging.ErrorLog(fmt.Sprintf("Unexpected signature method: %v.", token.Header["alg"]), nil)
			return nil, fmt.Errorf("unexpected signature method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil // 署名が正しければJWT_SECRET_KEYをバイト配列にして返す
	})
	if err != nil {
		return nil, err
	}

	// トークン自体が有効か秘密鍵を用いて確認。また、クレーム部分も取得。(トークンの署名が正しいか、有効期限内か、ブラックリストでないか。)
	claims, ok := token.Claims.(jwt.MapClaims) // MapClaimsにアサーション
	if !ok || !token.Valid {                   // 取得に失敗または検証が失敗
		return nil, errors.New("invalid authentication token")
	} else { // 有効な場合クレームを検証
		// idを検証
		id, ok := claims["id"].(int)
		if !ok {
			return nil, errors.New("id could not be obtained from the token")
		}
		if err := model.CfmId(id); err != nil { // ユーザーに存在するか。
			return nil, err
		}

		// expを検証
		exp, ok := claims["exp"].(int64)
		if !ok {
			return nil, errors.New("exp could not be obtained from the token")
		}
		expTT := time.Unix(exp, 0) // Unix 時刻を日時に変換
		timeNow := time.Now()      // 現在時刻を取得
		if timeNow.Before(expTT) { // エラーになるパターン  // 期限expTTが現在時刻timeNowより前ならtrue
			return nil, err
		}
	}

	// 正常に終われば解析されたトークンを渡す。
	return token, nil
}
