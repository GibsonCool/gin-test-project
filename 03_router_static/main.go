package main

import (
	"github.com/gin-gonic/gin"
	. "log"
	"net/http"
)

/*
	静态文件夹
*/
func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/test.JPG", "./test.JPG")

	/*
		运行时，需要在终端切刀 当前 main.go 文件目录下使用 go run 运行，否则上面的静态文件路径指向会找不到
	*/
	if err := r.Run(); err != nil {
		Println(err.Error())
	}

}
