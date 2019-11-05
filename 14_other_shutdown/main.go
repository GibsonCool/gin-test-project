package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

/*
	想要优雅地重启或停止你的Web服务器，使用下面的方法
    我们可以使用fvbock/endless来替换默认的ListenAndServe，
	https://segmentfault.com/a/1190000013757098

	go > 1.8 版本可以使用 http.Server 内置的 Shutdown() 方法来关闭如下
*/

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "welcome to Gin server")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	//  使用协程去执行服务，不阻塞
	go func() {
		if e := server.ListenAndServe(); e != nil && e != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", e.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	/*
		Go | http.Server 优雅退出方式简析 https://ictar.xyz/2018/09/04/Go-%E7%9A%84-http-Server-%E4%BC%98%E9%9B%85%E9%80%80%E5%87%BA%E6%96%B9%E5%BC%8F%E7%AE%80%E6%9E%90/
	*/
	if e := server.Shutdown(ctx); e != nil {
		log.Fatal("Server Shutdown :", e)
	}

}
