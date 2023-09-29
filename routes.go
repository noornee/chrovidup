package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1/videos")

	v1.POST("/upload", uploadVideo)
	v1.StaticFS("", http.Dir(temp_dir))

	// no route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found.",
			"status":  http.StatusNotFound,
		})
	})

	return r
}
