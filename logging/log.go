package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// ログファイル出力のセットアップ
func SetupLogging() (*os.File, error) {
	// ログファイルを作成
	logFile, err := os.OpenFile("./logging/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil { // エラーチェック
		return nil, fmt.Errorf("error opening file: %v", err) // エラーの場合
	}

	// ログの出力先をファイルにも。
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.Printf("Set up logging.\n\n")

	return logFile, nil // ファイルを返す
}

// エラー時のログをログファイルに残す
func ErrorLog(errName string, err error) {
	log.Printf("ERROR LOG: %s\n", errName)
	log.Printf("Time: %v\n", time.Now()) // 時刻
	log.Printf("Error: %s\n\n", err)
}
