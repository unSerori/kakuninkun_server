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
	// root page
	engine.GET("/", controller.ShowRootPage)
	// json test
	engine.GET("/test/json", controller.TestJson)
	// register user
	engine.POST("/users/register", controller.RegisterUser)

	return engine, nil // router設定されたengineを返す。
}
