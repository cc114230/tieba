package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

func CaptchaHandler(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.L().Error("生成验证码错误", zap.Error(err))
		return
	}
	ResponseSuccess(c, gin.H{
		"captchaID": id,
		"picPath":   b64s,
	})
}
