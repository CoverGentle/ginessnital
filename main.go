package main

import (
	"ginessnital/global"
	"ginessnital/router"
	"ginessnital/utils/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	global.Mysql = db.InitDB()
	r = router.CollectRoutes(r)

	r.Run(":8080")
}
