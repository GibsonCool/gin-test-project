package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 自定义中间件白名单函数
func IpAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
			"192.168.1.11",
		}
		flag := false

		for _, host := range ipList {
			if c.ClientIP() == host {
				flag = true
				break
			}
		}
		if !flag {
			c.String(http.StatusUnauthorized, " %s ，not in iplist", c.ClientIP())
			c.Abort()
		}
	}
}

func main() {

	r := gin.Default()

	// 使用白名单中间件
	r.Use(IpAuthMiddleware())

	/*
		curl -X GET "http://127.0.0.1:8080/testip"
	*/
	r.GET("/testip", func(c *gin.Context) {
		c.String(http.StatusOK, "hello testip")
	})

	r.Run()
}
