package dao

import (
	"fmt"
	"log"
	"net/http"
)

/**
 *@Description
 *@author shuaibo.tang
 *@Date 2022/11/23 14:44
 */

// 获取天气
func GetWeather() (res *http.Response, err error) {
	client := &http.Client{}
	url := "http://t.weather.sojson.com/api/weather/city/101281601"
	reqest, err := http.NewRequest("GET", url, nil)
	log.Println(&reqest, "reqest")
	if err != nil {
		fmt.Println(err)
	}
	resp, _ := client.Do(reqest)

	status := resp.Status
	fmt.Println(resp, "resp")
	fmt.Println("1111111111111111111111", status)
	return res, err
}

// 获取城市
func GetCityList() {
	//client := http.Client{}
}
