package route

import (
	"kakuninkun_server/controller"
	"kakuninkun_server/middleware"

	"github.com/gin-gonic/gin"
)

func GetRouter() (*gin.Engine, error) {
	// エンジンを作成
	engine := gin.Default()

	// MidLog all
	engine.Use(middleware.MidLog())

	// endpoints
	// root page
	engine.GET("/", controller.ShowRootPage)
	// json test
	engine.GET("/test/json", controller.TestJson)
	// cfm req
	engine.POST("/test/cfmreq", controller.CfmReq)
	// register user
	engine.POST("/users/register", controller.RegisterUser)

	return engine, nil // router設定されたengineを返す。
}
