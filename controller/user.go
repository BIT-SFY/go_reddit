package controller

import (
	"errors"
	"reddit/dao/mysql"
	"reddit/logic"
	"reddit/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// @Summary 注册
// @Description 处理用户注册请求
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param content body string true "json"
// @Success 1000 {object} _ResponseMessage
// @Router /api/v1/signup [post]
func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil { //只能校验字段类型对不对,
		//请求参数有误,直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok { //如果不是validator.ValidationErrors类型的错误
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

// @Summary 登录
// @Description 处理用户登录请求
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param content body string true "json"
// @Success 1000 {object} _ResponseMessage
// @Router /api/v1/login [post]
func LoginHandler(c *gin.Context) {
	//1.获取请求参数并检验参数
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.执行业务逻辑
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login(p) failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExit)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3.返回结果
	ResponseSuccess(c, gin.H{
		"user_id":   user.UserID,
		"user_name": user.Username,
		"token":     user.Token,
	})
}
