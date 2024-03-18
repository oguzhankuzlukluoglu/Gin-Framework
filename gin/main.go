package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("getData", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "hi ı am gin framework",
		})

	})

	router.Run()
}
