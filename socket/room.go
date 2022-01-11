package socket

import (
	"fmt"
	"vue-game/store"
)

type Room struct {
	ID    string
	Owner string
	Name  string
}

func (room Room) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":    room.ID,
		"owner": room.Owner,
		"name":  room.Name,
	}
}

func CreateRoom(user User) {
	redis := store.GetRedis()
	key := fmt.Sprintf("room:%s", user.Token)
	room := Room{
		ID:    "",
		Owner: "",
		Name:  "",
	}
	redis.HMSet(key, room.ToMap())
}

func GetUserRooms(userToken string) []Room {
	return nil
}
