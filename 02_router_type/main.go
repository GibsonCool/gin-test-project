package main

import (
	"github.com/gin-gonic/gin"
	. "log"
	"net/http"
)

/*
	gin 各种路由设置的方式
*/

func main() {

	r := gin.Default()

	r.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "get method"})
	})

	r.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "post method"})
	})

	// 上面的方式最终都是对调用这种方式
	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "handle delete method"})
	})

	// 内部各种方式的都创建一边路由，所以都支持
	r.Any("/any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "any method from " + c.Request.Method})
	})

	if err := r.Run(":8080"); err != nil {
		Println(err.Error())
	}
}
