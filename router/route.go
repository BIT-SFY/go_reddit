package router

import (
	"net/http"
	"reddit/controller"
	"reddit/logger"
	"reddit/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册
	r.POST("/signup", controller.SignUpHandler)
	//登录
	r.POST("/login", controller.LoginHandler)

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("version"))
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})
	return r
}
