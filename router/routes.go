package router

import (
	"ginessnital/controller"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/user/list", controller.CheckUsers)
	r.POST("/user/register", controller.RegisterUserInfo)
	r.POST("/user/login", controller.LoginUserInfo)
	return r
}
