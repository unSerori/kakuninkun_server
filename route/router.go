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
	engine.POST("/test/cfmreq", middleware.MidAuthToken(), controller.CfmReq)
	// register user
	engine.POST("/users/register", controller.RegisterUser)
	// user login
	engine.POST("/users/login", controller.Login)
	// get user data
	engine.GET("/users/user", middleware.MidAuthToken(), controller.UserProfile)
	// get users data
	engine.GET("/users/list", middleware.MidAuthToken(), controller.UsersDataList)
	// get companies and kgroups list
	engine.GET("/companies/list", controller.CompList)
	// update situation
	engine.POST("/users/situation", middleware.MidAuthToken(), controller.UpdateSitu)
	// cfm login
	engine.GET("/users/cfmlogin", middleware.MidAuthToken(), controller.Cfmlogin)
	// delete account
	engine.DELETE("/users/:id", middleware.MidAuthToken(), controller.DeleteUser)

	return engine, nil // router設定されたengineを返す。
}
