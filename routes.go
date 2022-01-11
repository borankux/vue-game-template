package main

import (
	"github.com/gin-gonic/gin"
	"vue-game/api"
)

func BindRoutes(app *gin.Engine) {
	app.GET("/api/rooms", api.GetRooms)
	app.POST("/api/room", api.CreateRoom)
}
