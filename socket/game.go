package socket

import (
	io "github.com/googollee/go-socket.io"
	"log"
)

func initGame(server *io.Server) {
	server.OnConnect("/game", func(conn io.Conn) error {
		log.Printf("game:connected:%s", conn.ID())
		recordIP(conn, "gamer")
		server.BroadcastToNamespace("/game", "user-update", getUsers())
		return nil
	})

	server.OnDisconnect("/game", func(conn io.Conn, s string) {
		defer func() {
			conn.Close()
		}()
		removeIP(conn, "gamer")
		log.Printf("game:disconnected:%s:%s", conn.ID(), s)
	})

	server.OnEvent("/game", "start", func(conn io.Conn, msg string) {
		log.Printf("/game started by:%s, with a message:%s", conn.ID(), msg)
	})

}
