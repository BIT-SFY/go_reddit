package controller

import (
	"reddit/logic"
	"reddit/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteController(c *gin.Context) {
	//参数请求校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) //翻译并去除掉错误提示结构体种的结构体标识
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	//获取当前请求的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	//投票的业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost(userID, p) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//返回结果
	ResponseSuccess(c, nil)
}
