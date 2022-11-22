package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 *@Description
 *@author shuaibo.tang
 *@Date 2022/11/22 14:45
 */

func AuthHandler(c *gin.Context) {
	fmt.Println("token鉴权")

}
