package room

import (
	"net"
)

type Room interface{
	UpdateUser(userId string,conn net.Conn,roomId string)
	DeleteUser(userId string)
	AddRoom(roomId string)
	RemoveRoom(roomId string)
} 
