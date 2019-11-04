package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	对绑定解析到结构体上的参数，做参数验证

	gin 默认是使用 go-playground/validator.v8 进行验证。
	查看标签用法的全部文档（http://godoc.org/gopkg.in/go-playground/validator.v8#hdr-Baked_In_Validators_and_Tags）.
*/
type Person struct {
	Age     int    `form:"age" binding:"required,gt=10"`
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()

	/*
		curl -X GET "http://127.0.0.1:8080/testing?age=34&name=ggtest&address=beijing"

		curl -X GET "http://127.0.0.1:8080/testing?age=10&name=ggtest&address=beijing"
	*/
	r.GET("/testing", func(c *gin.Context) {
		var person Person
		if e := c.ShouldBind(&person); e == nil {
			c.String(http.StatusOK, "%v", person)
		} else {
			c.String(http.StatusOK, "person bind err:%v", e.Error())
		}
	})
	r.Run()
}
