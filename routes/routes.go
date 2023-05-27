package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"liaoBa/controller"
	_ "liaoBa/docs"
	"liaoBa/logger"
	"liaoBa/middlewares"
	"net/http"
)

func SetUp(mode string) *gin.Engine {
	if mode == "realize" {
		gin.SetMode(gin.ReleaseMode) //gin设置为发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")

	//注册
	v1.POST("/signup", controller.SignUpHandler)
	//登录
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用认证中间件

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts/", controller.GetPostListHandler)
		// 根据时间或者分数，获取帖子列表
		v1.GET("/posts2/", controller.GetPostListHandler2)
		// 投票
		v1.POST("/vote", controller.PostVoteController)
	}

	//r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
	//	// 判断当前请求头中是否有有效的jwt，如果是登录的用户才可用这个功能
	//	//c.Request.Header.Get("Authorization")
	//	c.String(http.StatusOK, "pong")
	//})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
