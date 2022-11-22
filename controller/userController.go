package controller

import (
	"ginessnital/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	username := c.PostForm("username")
	password := c.PostForm("password")
	age := c.PostForm("age")
	log.Println(username, "username", password, "password", age, "age")
	ageC, _ := strconv.Atoi(age)
	err := dao.RegisterUser(username, password, ageC)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "msg": "该用户名已被注册"})
	} else {
		c.JSON(200, gin.H{"code": 2000, "msg": "注册成功"})
	}
}

// 用户登录
func LoginUserInfo(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	result := dao.LoginUser(username)
	if result != nil {
		if result.Username != username || result.Password != password {
			c.JSON(200, gin.H{"code": 4001, "data": result, "msg": "账号或密码错误"})
		} else {
			c.JSON(200, gin.H{"code": 2000, "data": result, "msg": "登录成功"})
		}
	} else {
		c.JSON(200, gin.H{"code": 4001, "data": result, "msg": "用户名不存在"})
	}

}
