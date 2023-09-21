package router

import (
	"net/http"
	"reddit/controller"
	"reddit/docs"
	"reddit/logger"
	"reddit/middlewares"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//引入gin-swagger渲染文档数据,添加swagger访问路由
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)
	// 优雅重启测试
	v1.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "test bro")
	})
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
		// 根据时间分数获取帖子列表
		v1.GET("/posts", controller.GetPostListHandler)
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
