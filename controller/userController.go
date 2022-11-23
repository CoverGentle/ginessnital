package controller

import (
	"ginessnital/dao"
	"ginessnital/middleware"
	"ginessnital/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckUsers(c *gin.Context) {
	users := dao.GetAllUsers()
	if users == nil {
		c.JSON(200, gin.H{"error": 1, "msg": "查询失败"})
	} else {
		c.JSON(200, gin.H{"error": 0, "msg": "查询成功", "data": users})
	}

}

// 注册用户
func RegisterUserInfo(c *gin.Context) {
	reg := model.User{}
	c.Bind(&reg)
	err := dao.RegisterUser(reg.Username, reg.Password, reg.Age)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "msg": "该用户名已被注册"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 2000, "msg": "注册成功"})
	}
}

// 用户登录
func LoginUserInfo(c *gin.Context) {
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	log := model.User{}
	c.Bind(&log)
	result := dao.LoginUser(log.Username)
	if result != nil {
		if result.Username != log.Username || result.Password != log.Password {
			c.JSON(200, gin.H{"code": 4001, "data": result, "msg": "账号或密码错误"})
		} else {
			tokenstring, _ := middleware.GetToken(log.Username)
			c.JSON(200, gin.H{"code": 2000, "token": "banner_" + tokenstring, "msg": "登录成功"})
		}
	} else {
		c.JSON(200, gin.H{"code": 4001, "data": result, "msg": "用户名不存在"})
	}

}
