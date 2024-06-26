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

	// endpoint
	// root page
	engine.GET("/", controller.ShowRootPage)
	// json test
	engine.GET("/test/json", controller.TestJson)

	// endpoints group
	// apiのグループ
	api := engine.Group("/api")
	{
		// ver1
		v1 := api.Group("v1")
		{
			// get companies and kgroups list
			v1.GET("/companies/list", controller.CompList)

			v1Users := v1.Group("/users")
			{
				// user login
				v1Users.POST("/login", controller.Login)
				// register user
				v1Users.POST("/register", controller.RegisterUser)
			}

			// 認証ミドルウェア適応グループ
			auth := v1.Group("/auth")
			auth.Use(middleware.MidAuthToken()) // 適応するミドルウェア
			{
				// cfm req
				auth.POST("/test/cfmreq", controller.CfmReq)

				// ユーザー関連グループ
				authUsers := auth.Group("/users")
				{
					// get user data
					authUsers.GET("/user", controller.UserProfile)
					// get users data
					authUsers.GET("/list", controller.UsersDataList)
					// update situation
					authUsers.POST("/situation", controller.UpdateSitu)
					// cfm login
					authUsers.GET("/cfmlogin", controller.Cfmlogin)
					// delete account
					authUsers.DELETE("/:id", controller.DeleteUser)
				}
			}
		}
	}

	// // root page
	// engine.GET("/", controller.ShowRootPage)
	// // json test
	// engine.GET("/test/json", controller.TestJson)
	// // cfm req
	// engine.POST("/test/cfmreq", middleware.MidAuthToken(), controller.CfmReq)
	// // register user
	// engine.POST("/users/register", controller.RegisterUser)
	// // user login
	// engine.POST("/users/login", controller.Login)
	// // get user data
	// engine.GET("/users/user", middleware.MidAuthToken(), controller.UserProfile)
	// // get users data
	// engine.GET("/users/list", middleware.MidAuthToken(), controller.UsersDataList)
	// // get companies and kgroups list
	// engine.GET("/companies/list", controller.CompList)
	// // update situation
	// engine.POST("/users/situation", middleware.MidAuthToken(), controller.UpdateSitu)
	// // cfm login
	// engine.GET("/users/cfmlogin", middleware.MidAuthToken(), controller.Cfmlogin)
	// // delete account
	// engine.DELETE("/users/:id", middleware.MidAuthToken(), controller.DeleteUser)

	return engine, nil // router設定されたengineを返す。
}
