package controller

import (
	"reddit/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ---和社区相关---

// CommunityHandler
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区,并以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed...", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不轻易把服务端具体的错误暴露在外卖
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler
func CommunityDetailHandler(c *gin.Context) {
	// 1.获取社区id
	idstr := c.Param("id") // 获取URL参数
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2.根据ID查询社区详情
	data, err := logic.GetCommunityListDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityListDetail() failed...", zap.Error(err))
		ResponseError(c, CodeInvalidCommunityID) //不轻易把服务端具体的错误暴露在外卖
		return
	}
	ResponseSuccess(c, data)
}
