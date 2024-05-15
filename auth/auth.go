package auth

import (
	"errors"
	"fmt"
	"kakuninkun_server/logging"
	"kakuninkun_server/model"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

	// uuid作成、登録処理
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	fmt.Println("uuid: " + uuid.String())
	if err := model.SaveJti(id, uuid.String()); err != nil {
		return "", err
	}

	// クレーム部分
	claims := jwt.MapClaims{
		"id":  id,            // id。
		"jti": uuid.String(), // uuid
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
func ParseToken(tokenString string) (*jwt.Token, int, error) {
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
		return nil, 0, err
	}

	// 下のクレーム検証処理(:elseスコープ内)で持ち出したい値をあらかじめ宣言しておく。
	var id int // id

	// トークン自体が有効か秘密鍵を用いて確認。また、クレーム部分も取得。(トークンの署名が正しいか、有効期限内か、ブラックリストでないか。)
	claims, ok := token.Claims.(jwt.MapClaims) // MapClaimsにアサーション
	if !ok || !token.Valid {                   // 取得に失敗または検証が失敗
		return nil, 0, errors.New("invalid authentication token")
	} else { // 有効な場合クレームの各要素を検証
		// idを検証
		idClaims, ok := claims["id"].(float64) // goではJSONの数値は少数もカバーしたfloatで解釈される
		if !ok {
			return nil, 0, errors.New("id could not be obtained from the token")
		}
		id = int(idClaims)                      // 調整してスコープ買いに持ち出す。
		if err := model.CfmId(id); err != nil { // ユーザーに存在するか。int(id)
			fmt.Print("idを検証ERR2: ")
			fmt.Println(err)

			return nil, 0, err
		}
		// expを検証
		exp, ok := claims["exp"].(float64)
		if !ok {
			return nil, 0, errors.New("exp could not be obtained from the token")
		}
		expTT := time.Unix(int64(exp), 0) // Unix 時刻を日時に変換
		timeNow := time.Now()             // 現在時刻を取得
		if timeNow.After(expTT) {         // エラーになるパターン  // 現在時刻timeNowが期限expTTより後ならエラーなのでtrueを出力
			return nil, 0, err
		}
	}

	// 正常に終われば解析されたトークンとidを渡す。
	return token, id, nil
}
