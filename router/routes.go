package router

import (
	"ginessnital/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/user/v1/check", controller.CheckUsers)
	r.POST("/user/register", controller.RegisterUserInfo)
	r.POST("/user/login", controller.LoginUserInfo)
	return r
}
