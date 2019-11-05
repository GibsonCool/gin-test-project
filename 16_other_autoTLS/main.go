package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": "001",
			"msg":  "ping 通过",
			"uri":  c.Request.RequestURI,
		})
	})
	/*
		大致流程是：
			1、生成本地的密钥
			2、将本地密钥发送给证书颁发机构获取一个 私钥
			3、对私钥进行验证，并保存起来
			4、下次请求 使用私钥加密
	*/
	autotls.Run(r, "www.itpp.tk")
}
