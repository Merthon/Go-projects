package routes

import (
	"github.com/Merthon/myginapp/controllers"
	"github.com/Merthon/myginapp/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 1. 加载 templates 目录下所有模板
	//r.LoadHTMLGlob("templates/*.tmpl")
	// 2. 提供 static 目录下的静态文件服务，对外映射到 /static
	//r.Static("/static", "/static")

	//r.GET("/", controllers.ShowIndex)
	//r := gin.New()
	//不在使用r := gin.Default()
	r.Use(gin.Recovery())       //捕获panic
	//使用自己定义的中间件
	r.Use(middlewares.Logger()) 
	//r.Use(middlewares.RateLimitMiddleware())
	//使用权限坚定AuthRequired
	/* api := r.Group("/api")
	{
		// // 登录接口无需鉴权
		// api.POST("/loginjson", controllers.LoginJSON)
		// api.GET("/ping", controllers.Ping)

		// // 需要鉴权的接口
		// auth := api.Group("/")
		// auth.Use(middlewares.AuthRequired())
		// {
		// 	auth.GET("/user/:id", controllers.GetUser)
		// 	auth.GET("/search", controllers.SearchItem)
		// 	auth.POST("/login", controllers.PostFormData)
		// }
		api.GET("/item/:id", controllers.GetItem)
        api.POST("/item/:id", controllers.PostItem)
	} */
	api := r.Group("/api")
	{
		api.GET("/users", controllers.ListUsers)
        api.GET("/users/:id", controllers.GetUser)
        api.POST("/users", controllers.CreateUser)
        api.PUT("/users/:id", controllers.UpdateUser)
        api.DELETE("/users/:id", controllers.DeleteUser)
	}
	return r
}
