package main

import (
	handlers "ginessnital/controller"
	"ginessnital/global"
	"ginessnital/utils/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	global.Mysql = db.InitDB()
	userv1_h := handlers.UserV1{}
	userv1 := r.Group("/user/v1")
	{
		userv1.GET("/check", userv1_h.CheckUsers)
	}
	r.Run(":8080")
}
