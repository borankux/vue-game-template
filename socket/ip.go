package socket

import (
	"fmt"
	io "github.com/googollee/go-socket.io"
	"vue-game/store"
	"vue-game/util"
)

func recordIP(conn io.Conn, key string) {
	ip, port := util.GetIPAndPort(conn.RemoteAddr())
	finalKey := fmt.Sprintf("%s:%s", key, ip)
	store.GetRedis().Set(finalKey, port, 0)
}

func removeIP(conn io.Conn, key string) {
	ip, _ := util.GetIPAndPort(conn.RemoteAddr())
	finalKey := fmt.Sprintf("%s:%s", key, ip)
	store.GetRedis().Del(finalKey)
}
