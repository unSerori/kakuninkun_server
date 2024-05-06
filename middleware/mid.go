package middleware

import (
	"kakuninkun_server/auth"
	"kakuninkun_server/logging"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ロギング
func MidLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// リクエストを受け取った時のログ
		log.Printf("Received request.\n")                        // リクエストの受理ログ
		log.Printf("Time: %v\n", time.Now())                     // 時刻
		log.Printf("Request method: %s\n", ctx.Request.Method)   // メソッドの種類
		log.Printf("Request path: %s\n\n", ctx.Request.URL.Path) // リクエストパラメータ

		// リクエストを次のハンドラに渡す
		ctx.Next()

		// レスポンスを返した後のログ
		log.Printf("Sent response.\n")                             // レスポンスの送信ログ
		log.Printf("Time: %v\n", time.Now())                       // 時刻
		log.Printf("Response Status: %d\n\n", ctx.Writer.Status()) // ステータスコード
	}
}

// トークン検証
func MidAuthToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ヘッダーからトークンを取得
		headerAuthorization := ctx.GetHeader("Authorization")
		if headerAuthorization == "" { // ヘッダーが存在しない場合
			// エラーログ
			logging.ErrorLog("Authentication unsuccessful.", nil)
			// レスポンス
			ctx.JSON(http.StatusBadRequest, gin.H{
				"srvResCode": 7001,                           // コード
				"srvResMsg":  "Authentication unsuccessful.", // メッセージ
				"srvResData": gin.H{},                        // データ
			})
			ctx.Abort() // 次のルーティングに進まないよう処理を止める。
			return      // 早期リターンで終了
		}

		// トークンの解析を行う。
		token, id, err := auth.ParseToken(headerAuthorization)
		if err != nil {
			// エラーログ
			logging.ErrorLog("Authentication unsuccessful. Maybe that user does not exist. Failed to parse token.", err)
			// レスポンス
			ctx.JSON(http.StatusBadRequest, gin.H{
				"srvResCode": 7008,                                                                                  // コード
				"srvResMsg":  "Authentication unsuccessful. Maybe that user does not exist. Failed to parse token.", // メッセージ
				"srvResData": gin.H{},                                                                               // データ
			})
			ctx.Abort() // 次のルーティングに進まないよう処理を止める。
			return      // 早期リターンで終了
		}

		ctx.Set("token", token) // トークンをコンテキストにセットする。  // _ = token // トークンを破棄。
		ctx.Set("id", id)       // 送信元クライアントのtokenのidを保持

		ctx.Next() // エンドポイントの処理に移行
	}
}
