// Custom http configuration with GIN
// Route grouping in GIN
// Using basic Auth functionally with GIN
package main

import (
	"github.com/gin-gonic/gin"
	"middleware/arakatman"
)

func main() {
	router := gin.New()
	//apply all routes	router.Use(arakatman.Authenticate)
	admin := router.Group("/admin", arakatman.Authenticate, arakatman.AddHeader)
	{
		admin.GET("/getData", getData)
		admin.GET("/getData1", getData1)

	}
	//	router.GET("/getData", arakatman.Authenticate, getData)
	//	router.GET("/getData1", getData1)
	router.GET("/getData2", getData2)
	router.Run()
}
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hi I am getData method",
	})
}

// http://localhost:8080/getQueryString?name=Mark&age=30

func getData1(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hi I am getData1 method",
	})
}
func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hi I am getData2 method",
	})
}
