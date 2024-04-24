package route

import (
	"Kakuninkun_server/controller"
	"Kakuninkun_server/middleware"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	// エンジンを作成
	engine := gin.Default()

	// endpoints
	// mid all
	engine.Use(middleware.MidLog)
	// /
	engine.GET("/", controller.ShowRootPage)
	// /json
	engine.GET("/json", controller.ShowTPage)

	return engine // router設定されたengineを返す。
}
