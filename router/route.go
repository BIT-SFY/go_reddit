package router

import (
	"net/http"
	"reddit/controller"
	"reddit/logger"

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

	r.GET("/ping", func(c *gin.Context) {
		// 如果是登录的用户
		c.String(http.StatusOK, viper.GetString("version"))
		// 否则就直接返回请登录
		c.String(http.StatusOK, "请登录")
	})
	return r
}
