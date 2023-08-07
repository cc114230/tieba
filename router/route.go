package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tieba/controller"
	"tieba/logger"
	"tieba/middleware"
)

func SetupRouter() *gin.Engine {

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.Use(middleware.Cors())
	v1 := r.Group("/api/v1")
	//注册路由
	v1.POST("/signup", controller.SignUpHandler)
	//登录路由
	v1.POST("/login", controller.LoginHandler)

	v1.GET("/community", controller.CommunityHandler)
	v1.GET("/community/:id", controller.CommunityDetailHandler)
	v1.GET("/posts", controller.GetPostListHandler)
	//按照帖子发布时间或点赞数获取帖子列表
	v1.GET("/posts2", controller.GetPostListHandler2)
	v1.GET("/post/:id", controller.GetPostDetailHandler)

	//v1.GET("/captcha", controller.CaptchaHandler)
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.POST("/post", controller.CreatePostHandler)

		v1.POST("/vote", controller.PostVoteHandler)
		v1.POST("/comment", controller.CommentHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r

}
