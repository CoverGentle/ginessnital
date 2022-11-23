package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

/**
 *@Description
 *@author shuaibo.tang
 *@Date 2022/11/22 14:45
 */

type CustomClaims struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        //内嵌声明标准
}

const TokenExpireDuration = time.Hour * 24

// CustomSecret 用于加盐的字符串
var CustomSecret = []byte("第一个go系统")

func GetToken(username string) (string, error) {
	claims := CustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "my-project", //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(CustomSecret)
}
func AuthCheck(username string, password string, c *gin.Context) {
	//var user model.User
	//err := c.ShouldBind(&user)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{"code": 2001, "msg": "无效的参数"})
	//	return
	//}
	//if user.Username == username && user.Password == password {
	//	tokenString = GetToken(username)
	//	c.JSON(http.StatusOK, gin.H{"code": 2000, "msg": "success", "data": gin.H{"token":}})
	//
	//}

	//fmt.Println("token鉴权")

}
