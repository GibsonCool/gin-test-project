package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {

	pwdPath, _ := os.Getwd()
	pwdPath += "/12_middleware_gin"
	// 创建两个日志文件，
	fLog, _ := os.Create(pwdPath + "/gin.log")
	fErrLog, _ := os.Create(pwdPath + "/gin_err.log")

	// 用创建的日志文件 writer 覆盖默认gin的 writer。这样 Logger() 中间件在输出的时候会使用gin 默认的writer也就是我们覆盖的这个wirter
	gin.DefaultWriter = io.MultiWriter(fLog)
	// 同样的错误日志文件 writeer 覆盖，使用 Recovery()中间件的时候也会将日志写入我们覆盖的writer中去
	gin.DefaultErrorWriter = io.MultiWriter(fErrLog)

	/*
		gin.Default()。默认使用 Logger() Recovery()
		Logger()：控制台输出的日志
		Recovery():异常的捕获，也输出到控制台，这样既避免服务直接宕机

		gin.New()不适用默认的两个中间件，
	*/
	r := gin.New()
	// 使用中间件
	r.Use(gin.Logger(), gin.Recovery())

	/*
		curl -X GET "http://127.0.0.1:8080/test/middleware?name=err"

		curl -X GET "http://127.0.0.1:8080/test/middleware?name=test"

		curl -X GET "http://127.0.0.1:8080/test/middleware?name=testing"
	*/
	r.GET("/test/middleware", func(c *gin.Context) {
		name := c.DefaultQuery("name", "default_test")
		// 如果符合条件出发错误，经过 Recovery中间件会将日志写入我们设置的 gin_err.log中，并不会导致程序崩溃
		if name == "err" {
			panic("test panic")
		}
		c.String(http.StatusOK, "name:%s", name)
	})

	r.Run()
}
