package controllers

import (
	"net/http"
	"strconv"

	"github.com/solnsumei/simple-chat/utils"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/simple-chat/models"
)

// GetOpenChat from database
func GetOpenChat(c *gin.Context) {
	authID := c.MustGet("authID").(int64)

	var chats []models.Chat
	models.DB.Where("sender_id = ?", authID).Or(
		"receiver_id = ?", authID).Find(&chats)

	c.JSON(http.StatusOK, gin.H{"chats": chats})
}

// GetChatMessages from database
func GetChatMessages(c *gin.Context) {
	authID := c.MustGet("authID").(int64)
	// Check if user Id passed in is valid
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		utils.BadRequestResponse(c, "User Id is invalid.")
		return
	}

	// fmt.Printf("%d, %d", authID, userID)

	if authID == int64(userID) {
		utils.BadRequestResponse(c, "You cannot get message from yourself.")
		return
	}

	var chat models.Chat
	if err := models.DB.Where(
		"sender_id = ? AND receiver_id = ?", userID, authID).Or(
		"sender_id = ? AND receiver_id = ?", authID, userID).First(&chat).Error; err != nil {

		chat.SenderID = uint(authID)
		chat.ReceiverID = uint(userID)

		if createErr := models.DB.Create(&chat).Error; createErr != nil {
			c.JSON(http.StatusOK, gin.H{"chat": nil, "messages": nil})
			return
		}
	}

	var messages []models.Message
	models.DB.Where("chat_id = ?", chat.ID).Find(&messages)

	c.JSON(http.StatusOK, gin.H{"chat": chat, "messages": messages})
}
