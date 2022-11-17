package handlers

import (
	"ginessnital/dao"
	"github.com/gin-gonic/gin"
)

type UserV1 struct{}

func (UserV1) CheckUsers(c *gin.Context) {
	users := dao.GetAllUsers()
	if users == nil {
		c.JSON(200, gin.H{"error": 1, "msg": "查询失败"})
	} else {
		c.JSON(200, gin.H{"error": 0, "msg": "查询成功", "users": users})
	}

}
