package room

import (
	"fmt"
	"net"
)

type RoomRepository struct{
	Conn   map[string]net.Conn
    UserId []string
}

type RoomService struct{
	Room  map[string]*RoomRepository    
}

func NewRoomService(room *RoomRepository) *RoomService{
	roomservice:=make(map[string]*RoomRepository)
	fmt.Print("map initialized")
	return &RoomService{
      Room : roomservice,
	}
}

func (r *RoomService) AddRoom(roomId string){
	_, prs := r.Room[roomId]

    fmt.Println("prs:", prs)

	if !prs{
		r.Room[roomId]=&RoomRepository{
			Conn: make(map[string]net.Conn),
		   }
	}
}

func (r *RoomService) RemoveRoom(roomId string){
  delete(r.Room,roomId)
}



func (r *RoomService) UpdateUser(userId string,conn net.Conn,roomId string) {
   r.Room[roomId].Conn[userId]=conn
   var ifPresent= false
   for _,value := range r.Room[roomId].UserId {
	if value==userId {ifPresent=true}
   }
   if ifPresent{
	 r.Room[roomId].UserId = append(r.Room[roomId].UserId, userId)
   }
}




func (r *RoomService) DeleteUser(userId string, roomId string) {
	delete(r.Room[roomId].Conn, userId)
	var index=-1
	for i,value := range r.Room[roomId].UserId {
	 if value==userId {index=i}
	}
	if index!=-1 {
 
	 userIds := append(r.Room[roomId].UserId[:index], r.Room[roomId].UserId[index+1:]...)
	 r.Room[roomId].UserId = userIds 
	}
 }



