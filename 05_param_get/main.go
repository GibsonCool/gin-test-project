package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*

 */
func main() {
	r := gin.Default()

	/*
		http://localhost:8080/test?first_name=wang
		http://localhost:8080/test?first_name=wang&last_name=qing
	*/
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "default_name")

		c.String(http.StatusOK, "%s  %s", firstName, lastName)
	})

	r.Run()
}
