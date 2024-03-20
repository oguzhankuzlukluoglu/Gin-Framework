package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getData", getData)
	router.GET("/getQueryString", getQueryString)
	router.POST("/getDataPost", getDataPost)
	router.GET("/getUrlData/:name/:age", getUrlData)
	router.Run()
}
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hi I am GET method from GIN framework",
	})

}
func getDataPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hi I am POST method from GIN framework",
	})
}

// http://localhost:8080/getQueryString?name=Mark&age=30

func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"data": "hi I am getQueryStringMethod",
		"name": name,
		"age":  age,
	})
}
func getUrlData(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"data": "Hi I am getUrlData method",
		"name": name,
		"age":  age,
	})
}
