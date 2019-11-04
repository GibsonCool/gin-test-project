package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	/*
		curl -d "first_name=faker&last_name=donib" http://127.0.0.1:8080/test
	*/
	r.POST("/test", func(c *gin.Context) {
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default_name")

		c.String(http.StatusOK, "%s  %s", firstName, lastName)
	})

	r.Run()
}
