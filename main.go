package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"vue-game/middleware"
	"vue-game/socket"
	"vue-game/store"
)

func main() {
	store.InitRedis()
	app := gin.Default()
	app.Use(middleware.Cors())
	socket.BindSockets(app)
	BindUI(app)
	BindRoutes(app)
	err := app.Run(":8000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
