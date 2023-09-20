package controller

import (
	"reddit/logic"
	"reddit/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	// 1.获取参数及参数的校验
	// c.ShouldBindJSON() // validator --> binding tag
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON failed", zap.Any("err", err))
		zap.L().Error("c.ShouldBindJSON failed")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从 c 取到当前发请求的用户ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情
func GetPostDetailHandler(c *gin.Context) {
	// 1.获取参数(帖子的id)
	pidStr := c.Param("id") //这个地方要和路由那个地方写的值一样
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.根据id去取出帖子数据(查数据库)
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表
func GetPostListHandler(c *gin.Context) {
	// 1.获取分页参数
	page, size, err := getPageInfo(c)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error(" logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版
// 根绝前端传来的参数动态获取帖子列表
// 按创建时间排序,按分数排序
// 1.获取参数
// 2.去redis查询id列表
// 3.根据id去数据库查询帖子详情
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数:/api/v1/posts2?page=1&size=10&order=time  ?后面的叫Query string参数,所以获取的时候都是c.Query这种方式
	// 1.初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
	}
	// c.ShouldBind() 动态的选择相应的数据类型获取数据
	// c.ShouldBindJSON() 如果请求中携带的是json格式的数据,才能用这个方法获取到数据

	// 2.获取数据
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error(" logic.GetPostList2() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, data)
}
