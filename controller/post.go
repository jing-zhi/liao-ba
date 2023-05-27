package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"liaoBa/logic"
	"liaoBa/models"
	"strconv"
)

// CreatePostHandler 创建帖子处理函数
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数，进行参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post invaild param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 从 c 获取到当前发请求的用户ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context) {
	// 从url获取帖子id
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 根据id取出帖子数据
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostByID(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表的处理函数
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	// 返回响应
}

// GetPostListHandler2 升级版,根据前端传来的参数，动态获取，按照分数或者创建时间
//@Summary升级版帖子列表接口
//@Description可按社区按时间或分数排序查询帖子列表接口
//@Tags帖子相关接口(api分组展示使用的)
//@Accept application/json
//@Produce application/json
//@Param Authorization header string true "Bearer JWT"
//@Param object query models.ParamPostList false "查询参数"
//@Security ApiKeyAuth
//@Success 200 {object} _ResponsePostList
//@Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数(query string): /api/v1/posts2?page=1&size=20&order=time
	// 获取分页参数
	// 初始化结构体时候指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // 不要 magic string
	}
	// c.ShouldBindQuery动态根绝类型获取参数
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invaild param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostListNew(p)
	//page, size := getPageInfo(c)
	// 获取数据
	if err != nil {
		zap.L().Error("logic.GetPostListNew() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	// 返回响应
}

// 1.获取请求的query string参数
// 2.去redis查询id列表
// 3.根据id去数据库查询帖子详细信息

//func GetCommunityPostListHandler(c *gin.Context) {
//	// 初始化结构体时候指定初始参数
//	p := &models.ParamCommunityPostList{
//		ParamPostList: models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//	}
//	// c.ShouldBindQuery动态根绝类型获取参数
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("GetCommunityPostListHandler with invaild param", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	//page, size := getPageInfo(c)
//	// 获取数据
//	data, err := logic.GetCommunityPostList(p)
//	if err != nil {
//		zap.L().Error("logic.GetCommunityPostList2() failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	ResponseSuccess(c, data)
//	// 返回响应
//}
