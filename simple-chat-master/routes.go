package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/solnsumei/simple-chat/controllers"
	"github.com/solnsumei/simple-chat/middlewares"
)

func loadGuestRoutes(router *gin.Engine) {
	guestRouter := router.Group("/api/v1")

	guestRouter.POST("/login", controllers.LoginUser)
	guestRouter.POST("/signup", controllers.RegisterUser)

	guestRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to simple chat api"})
	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to simple chat!"})
	})
}

func loadAuthRoutes(router *gin.Engine) {
	authRouter := router.Group("/api/v1")
	authRouter.Use(middlewares.AuthMiddleware())

	authRouter.GET("/user", controllers.GetProfile)
	authRouter.GET("/users", controllers.FetchAllUsers)
	authRouter.GET("/users/:userID", controllers.GetUser)
	authRouter.GET("/users/:userID/chat", controllers.GetChatMessages)
	authRouter.GET("/chats", controllers.GetOpenChat)
}

func socketHandler(router *gin.Engine) {
	socketRouter := router.Group("/")
	socketRouter.Use(middlewares.AuthMiddleware())

	socketRouter.GET("/socket.io/*any", gin.WrapH(controllers.SocketServer))
	socketRouter.POST("/socket.io/*any", gin.WrapH(controllers.SocketServer))
}
