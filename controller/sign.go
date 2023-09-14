package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tieba/dao/redis"
	"tieba/logic"
)

func SignHandler(c *gin.Context) {
	// 从 c 取到当前发请求的用户的ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	if err := logic.Sign(userID); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		if errors.Is(err, redis.ErrIsAlreadyCheckin) {
			ResponseError(c, CodeIsAlreadyCheckIn)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
