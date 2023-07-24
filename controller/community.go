package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"tieba/logic"
)

func CommunityHandler(c *gin.Context) {
	//拿到社区community_id和community_name返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区id
	idStr := c.Param("id") // 获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 根据id获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
