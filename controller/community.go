package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"liaoBa/logic"
	"strconv"
)

// 社区相关

func CommunityHandler(c *gin.Context) {
	// 查询到所有社区的 community_id community_name 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易向外暴露具体错误
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	// 1。 获取社区id
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 查询到所有社区的 community_id community_name 以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易向外暴露具体错误
		return
	}
	ResponseSuccess(c, data)
}
