package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var openAPISpec = gin.H{
	"openapi": "3.0.3",
	"info": gin.H{
		"title":       "MetLab10 Gin OpenAPI Demo",
		"version":     "1.0.0",
		"description": "Demo API with explicit OpenAPI endpoint.",
	},
	"paths": gin.H{
		"/ping": gin.H{
			"get": gin.H{
				"summary":     "Health check",
				"operationId": "ping",
				"responses": gin.H{
					"200": gin.H{
						"description": "OK",
						"content": gin.H{
							"application/json": gin.H{
								"schema": gin.H{
									"type": "object",
									"properties": gin.H{
										"ok": gin.H{"type": "boolean"},
									},
									"required": []string{"ok"},
								},
							},
						},
					},
				},
			},
		},
	},
}

func newRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	r.GET("/openapi.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, openAPISpec)
	})

	return r
}

func main() {
	r := newRouter()
	_ = r.Run("127.0.0.1:8101")
}

