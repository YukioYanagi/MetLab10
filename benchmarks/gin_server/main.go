package main

import "github.com/gin-gonic/gin"

func newRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	static := []byte(`{"ok":true}`)

	r.GET("/ping", func(c *gin.Context) {
		// Отдаём статические байты, без дополнительных сериализаций.
		c.Data(200, "application/json; charset=utf-8", static)
	})

	return r
}

func main() {
	r := newRouter()

	// Отличаем порт от FastAPI, чтобы можно было запускать параллельно.
	_ = r.Run("127.0.0.1:8001")
}

