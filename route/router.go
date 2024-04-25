package route

import (
	"kakuninkun_server/controller"
	"kakuninkun_server/middleware"

	"github.com/gin-gonic/gin"
)

func GetRouter() (*gin.Engine, error) {
	// エンジンを作成
	engine := gin.Default()

	// endpoints
	// MidLog all
	engine.Use(middleware.MidLog())
	// /
	engine.GET("/", controller.ShowRootPage)
	// /json
	engine.GET("/json", controller.ShowTPage)

	return engine, nil // router設定されたengineを返す。
}
