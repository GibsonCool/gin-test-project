package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	这里不能直接使用 idea 的直接运行 ，否则 html 目录找不到
	需要切换到文件目录下 终端执行   go run main.go
*/
func main() {

	r := gin.Default()

	// 设置 html 目录
	r.LoadHTMLGlob("template/*")
	// 也可以单独加载某个 HTML 文件
	//r.LoadHTMLFiles("template/index1.html","template/index2.html")

	/*
		curl -X GET "http://127.0.0.1:8080/index1"

		curl -X GET "http://127.0.0.1:8080/index2"
	*/
	r.GET("/index1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index1.html", gin.H{
			"title": "我的首页",
			"msg":   "来自模板渲染的内容",
		})
	})

	r.GET("/index2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index2.html", gin.H{
			"title": "我的首页",
			"msg":   "来自模板渲染的内容 index2 ----tt",
		})
	})

	r.Run()
}
