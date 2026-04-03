package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type loggerLike interface {
	Printf(format string, v ...any)
}

func requestLoggingMiddleware(logger loggerLike) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		elapsed := time.Since(start)

		logger.Printf(
			"method=%s path=%s status=%d duration_ms=%d",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			elapsed.Milliseconds(),
		)
	}
}

func newRouter(logger loggerLike) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(requestLoggingMiddleware(logger))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	return r
}

func main() {
	logger := log.New(os.Stdout, "app ", log.LstdFlags)
	r := newRouter(logger)
	_ = r.Run("127.0.0.1:8201")
	fmt.Println("server stopped")
}

