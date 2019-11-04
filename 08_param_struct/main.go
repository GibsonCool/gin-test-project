package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
	参数绑定到结构体来解析
*/
func main() {
	r := gin.Default()

	/*
		curl -X GET "http://127.0.0.1:8080/testing?name=adz&address=tianjin"
		curl -X GET "http://127.0.0.1:8080/testing?name=adz&address=tianjin&birthday=2019-11-11"

		curl -X POST "http://127.0.0.1:8080/testing?name=adz&address=tianjin&birthday=2019-11-11"
		curl -X POST "http://127.0.0.1:8080/testing" -d "name=adz&address=tianjin&birthday=2019-11-11"

		curl -H "Content-Type:application/json" -X POST "http://127.0.0.1:8080/testing" -d '{"name":"yitian","address":"beijing"}'
	*/
	r.GET("/testing", handlerTesting)
	r.POST("/testing", handlerTesting)

	r.Run()
}

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func handlerTesting(c *gin.Context) {
	var person Person
	/*
		这里会根据 content-type 去获取对应的 Binding 去做不通的处理
		所以不论是 json  form 等等参数都能绑定解析
	*/
	if e := c.ShouldBind(&person); e == nil {
		c.String(http.StatusOK, "%v", person)
	} else {
		c.String(http.StatusOK, "person bind err:%v", e.Error())
	}

}
