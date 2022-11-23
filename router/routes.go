package router

import (
	"ginessnital/controller"
	"ginessnital/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	r.GET("/user/list", controller.CheckUsers)
	r.POST("/user/register", controller.RegisterUserInfo)
	r.POST("/user/login", controller.LoginUserInfo)
	r.GET("/weather/info", controller.GetWeatherInfo)
	return r
}
