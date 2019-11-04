package main

import (
	"github.com/gin-gonic/gin"
	. "log"
	"net/http"
)

/*
	参数作为URL------》获取URL中的参数
*/
func main() {
	r := gin.Default()

	// http://localhost:8080/test/faker/123
	r.GET("/test/:name/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": c.Param("name"),
			"id":   c.Param("id"),
			"msg":  c.Request.RequestURI,
		})
	})

	// 上面的路由和这条路由规则一样，只能全路径匹配，
	// 这个路由规则能够匹配 /user/double 这种格式，但不能匹配 /user/ 或  /user 这种 URL 格式
	// http://localhost:8080/user/faker
	r.GET("/user/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": c.Param("name"),
			"msg":  c.Request.RequestURI,
		})
	})

	/*
		泛绑定：
			这个规则既能匹配 /file/java/ 格式也能匹配 /file/java/send 这种格式
			http://localhost:8080/file/java/
			http://localhost:8080/file/java/dotest
			http://localhost:8080/file/java/dotest/doadmin/test
	*/

	r.GET("/file/:filename/*action", func(c *gin.Context) {
		name := c.Param("filename")
		action := c.Param("action")
		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
			"msg":    c.Request.RequestURI,
		})
	})

	if err := r.Run(); err != nil {
		Println(err.Error())
	}

}
