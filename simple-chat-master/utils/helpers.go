package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFoundResponse return
func NotFoundResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found."})
}

// ForbiddenResponse return
func ForbiddenResponse(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "You are not authroized to view this resource."})
}

// BadRequestResponse return
func BadRequestResponse(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, gin.H{"error": message})
}
