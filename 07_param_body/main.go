package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()

	/*
		curl -X POST http://127.0.0.1:8080/test  -d "first_name=faker&last_name=donib"
		curl -X POST http://127.0.0.1:8080/test  -d {"name":"wagn"}
	*/
	r.POST("/test", func(c *gin.Context) {
		bodyBytes, e := ioutil.ReadAll(c.Request.Body)
		if e != nil {
			c.String(http.StatusBadRequest, e.Error())
			c.Abort()
		}

		/*
			如果想要在解析参数，需要把读完的数据从新写会 body 中
		*/
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default_name")

		c.String(http.StatusOK, "%s  %s  %s", firstName, lastName, string(bodyBytes))
	})

	r.Run()
}
