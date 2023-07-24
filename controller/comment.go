package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"tieba/logic"
	"tieba/models"
)

func CommentHandler(c *gin.Context) {
	//1.获取参数校验
	p := new(models.Comment)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Comment with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.业务处理
	// 从 c 取到当前发请求的用户的ID和name
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.CommenterID = userID

	Username, err := getCurrentUserName(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.Commenter = Username

	//发布评论
	if err = logic.Comment(p); err != nil {
		zap.L().Error("logic.Comment(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}
