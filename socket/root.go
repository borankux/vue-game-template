package socket

import (
	io "github.com/googollee/go-socket.io"
	"log"
	"time"
)

const HeartBeat = time.Second * 1

func initRoot(server *io.Server) {
	server.OnConnect("/", func(conn io.Conn) error {
		log.Printf("root:connected:%s\n", conn.ID())
		recordIP(conn, "admin")
		return nil
	})

	server.OnDisconnect("/", func(conn io.Conn, s string) {
		defer func() {
			conn.Close()
		}()
		removeIP(conn, "admin")
		log.Printf("disconnected:%s, %s\n", s, conn.ID())
	})

	server.OnError("/", func(s io.Conn, e error) {
		log.Println("meet error:", e)
	})

	go func() {
		for {
			server.BroadcastToNamespace("/", "stats", getStats(server))
			time.Sleep(HeartBeat)
		}
	}()
}
