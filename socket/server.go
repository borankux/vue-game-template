package socket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	io "github.com/googollee/go-socket.io"
)

func BindSockets(engine *gin.Engine) {
	server := io.NewServer(nil)

	initRoot(server)
	initGame(server)

	go func() {
		err := server.Serve()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
	engine.GET("/socket.io/*any", gin.WrapH(server))
	engine.POST("/socket.io/*any", gin.WrapH(server))
	engine.GET("/game/*any", gin.WrapH(server))
	engine.POST("/game/*any", gin.WrapH(server))
}
