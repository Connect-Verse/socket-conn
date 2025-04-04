package main

import (
	"log"
	"net/http"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/rs/cors"
	"github.com/saransh-g1/socket-conn/internal/room"
)



func main() {
	mux := http.NewServeMux()

    roomRepo:= room.RoomRepository{}
	roomService:=room.NewRoomService(&roomRepo)
	//rdb:= utils.CreateRedis()
    
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    
    userId := r.URL.Query().Get("userId")
	roomId := r.URL.Query().Get("roomId")

		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Print(err,"sjfaowijep")
			return
		}

		roomService.AddRoom(roomId)
		roomService.UpdateUser(userId,conn,roomId)


		defer func(){
            roomService.DeleteUser(userId,roomId)
			conn.Close()
		}()

			for {
				
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					log.Print(err,"sjfaowijep")
					break
				}

			for _, value:= range roomService.Room[roomId].Conn{
				err = wsutil.WriteServerMessage(value, op, msg)
				if err != nil {
					log.Print(err,"sjfaowijep")
					break
				}
			}
			}
		
	})

	// Apply CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)

	log.Println("WebSocket server listening on :8000")
	http.ListenAndServe(":8000", handler)
}