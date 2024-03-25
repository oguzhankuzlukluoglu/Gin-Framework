package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/simple-chat/utils"
)

func unauthorizedResponse(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": message})
	c.Abort()
}

// AuthMiddleware for validating logged in users
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		jwtToken := c.GetHeader("Authorization")
		if jwtToken == "" {
			queryToken := c.Query("token")
			if queryToken == "" {
				unauthorizedResponse(c, "Unauthorized access, please login.")
				return
			}
			jwtToken = "Bearer " + queryToken
		}

		// Split token string
		extractedToken := strings.Split(jwtToken, "Bearer ")
		if len(extractedToken) != 2 {
			unauthorizedResponse(c, "Unauthorized access, invalid token.")
			return
		}

		// Extract token string from it
		authToken := strings.TrimSpace(extractedToken[1])
		userID, err := utils.CheckAndVerifyToken(authToken)

		// If userID is not in claims return error
		if err != nil {
			unauthorizedResponse(c, "Unauthorized access, invalid token.")
			return
		}

		c.Set("authID", userID)
		c.Next()
	}
}
