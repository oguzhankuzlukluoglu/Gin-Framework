package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {

	f, _ := os.Create("ginLogging.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()
	router.GET("/getData", getData)
	router.Run()
}
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hi i am getData method",
	})
}
