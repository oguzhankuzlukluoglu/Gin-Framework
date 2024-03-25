package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/simple-chat/models"
	"github.com/solnsumei/simple-chat/utils"
)

// RegisterUser to database
func RegisterUser(c *gin.Context) {
	var input utils.UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := input.ValidateSignUp(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Hash password
	hashedPassword, hashingErr := utils.HashPassword(input.Password)
	if hashingErr != nil {
		log.Println(hashingErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": hashingErr})
		return
	}

	// Create user
	user := models.User{Name: input.Name, Email: input.Email, Password: hashedPassword}

	// Check errors in creating user
	if err := models.DB.Create(&user).Error; err != nil {
		log.Println(err)
		c.JSON(http.StatusConflict, gin.H{"error": "Email has already been taken"})
		return
	}

	token, err := utils.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Token could not be created, please try again later."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user, "token": token})
}

// LoginUser to app
func LoginUser(c *gin.Context) {
	var input utils.UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := input.ValidateLogin(); err != nil {
		log.Println(">>>>>", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var user models.User

	// fetch user from db
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and/or password is incorrect."})
		return
	}

	// Compare passwords
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and/or password is incorrect."})
		return
	}

	token, err := utils.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Token could not be created, please try again later."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}
