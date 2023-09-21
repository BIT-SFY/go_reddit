package router

import (
	"net/http"
	"reddit/controller"
	"reddit/logger"
	"reddit/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)
	// 只有用户登录才可访问以下路由
	v1.Use(middlewares.JWTAuthMiddleware()) //JWT认证中间件
	{
		// 获取所有的社区信息
		v1.GET("/community", controller.CommunityHandler)
		// 根据社区id获取社区信息
		v1.GET("/community/:id", controller.CommunityDetailHandler) //路径参数:id
		// 发布一个帖子
		v1.POST("/post", controller.CreatePostHandler)
		// 根据帖子id获取帖子信息
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		// 直接获取所有列表
		v1.GET("/posts", controller.GetPostListHandler)
		// 根据时间分数获取帖子列表
		v1.GET("/posts2", controller.GetPostListHandler2)
		// 投票
		v1.POST("/vote", controller.PostVoteController)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})
	return r
}
