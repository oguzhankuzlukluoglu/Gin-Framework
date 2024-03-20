package arakatman

import "github.com/gin-gonic/gin"

func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("Token") == "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"message": "Token Not present",
		})
		return
	}
	c.Next()
}
func Authenticate1() gin.HandlerFunc { //buradaki farklı bir authenticate yapısı, genelde bu kullanılır.
	//write custom logic to be applied before my middleware is executed
	return func(c *gin.Context) {
		if !(c.Request.Header.Get("Token") == "auth") {
			c.AbortWithStatusJSON(500, gin.H{
				"message": "Token Not present",
			})
			return
		}
		c.Next()
	}
}
func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("key", "value")
	c.Next()
}
