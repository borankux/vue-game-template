package socket

import (
	socketio "github.com/googollee/go-socket.io"
	"vue-game/store"
)

type Stats struct {
	Admins []string `json:"admins"`
	Users  []string `json:"users"`
	Rooms  []string `json:"rooms"`
}

func getAdmins() []string {
	slices := store.GetRedis().Keys("admin:*")
	var users []string
	slices.ScanSlice(&users)
	return users
}

func getUsers() []string {
	var users []string
	store.GetRedis().Keys("gamer:*").ScanSlice(&users)
	return users
}

func getStats(server *socketio.Server) Stats {
	return Stats{
		Admins: getAdmins(),
		Users:  getUsers(),
		Rooms:  server.Rooms("game"),
	}
}
