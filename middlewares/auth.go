package middlewares

import (
	"reddit/controller"
	"reddit/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxxx.xxxx.xxx

		// 获取Authorization
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" { //请求头中没有jwt
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort() //直接返回响应，不会进行下一步
			return
		}
		// 按空格分割Authorization
		parts := strings.SplitN(authHeader, " ", 2) //最多返回2个子字符串，最后一个字符串为未分割的剩余字符串
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，使用解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将解析JWT得到的用户ID,保存到请求的上下文c上
		// 后续的处理请求的函数可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
		c.Set(controller.CtxUserIDKey, mc.UserID)
		//继续进行下一步操作
		c.Next()
	}
}
