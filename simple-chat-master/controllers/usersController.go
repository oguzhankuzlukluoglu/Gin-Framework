package controllers

import (
	"net/http"
	"strconv"

	"github.com/solnsumei/simple-chat/utils"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/simple-chat/models"
)

// FetchAllUsers from database
func FetchAllUsers(c *gin.Context) {
	authID := c.MustGet("authID").(int64)

	var users []models.User
	models.DB.Where("ID <> ?", authID).Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GetProfile from database
func GetProfile(c *gin.Context) {
	authID := c.MustGet("authID").(int64)

	var user models.User
	if err := models.DB.Where("ID = ?", authID).First(&user).Error; err != nil {
		utils.NotFoundResponse(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GetUser profile from database
func GetUser(c *gin.Context) {
	// Check if user Id passed in is valid
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		utils.BadRequestResponse(c, "User Id is invalid.")
		return
	}

	var user models.User
	if findErr := models.DB.Where("ID = ?", userID).First(&user).Error; findErr != nil {
		utils.NotFoundResponse(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
