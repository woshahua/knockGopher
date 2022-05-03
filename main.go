package main

import "github.com/gin-gonic/gin"
import "runtime"

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	router.GET("/os", func(c *gin.Context) {
		c.String(200, runtime.GOOS)
	})
	router.Run(":8000")
}
