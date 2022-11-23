package controller

import (
	"ginessnital/dao"
	"github.com/gin-gonic/gin"
)

/**
 *@Description
 *@author shuaibo.tang
 *@Date 2022/11/23 17:21
 */

func GetWeatherInfo(c *gin.Context) {
	res, err := dao.GetWeather()

	c.JSON(200, gin.H{"data": res, "error": err})

}
