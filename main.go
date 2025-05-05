package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/rs/cors"
	// "github.com/saransh-g1/socket-conn/internal/grpc/client"
	"github.com/saransh-g1/socket-conn/internal/pub-sub"
	"github.com/saransh-g1/socket-conn/internal/room"
	"github.com/saransh-g1/socket-conn/internal/utils"
	// pb "github.com/saransh-g1/socket-conn/internal/grpc"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)



func main() {

	ctx := context.Background()

	mux := http.NewServeMux()

    roomRepo:= room.RoomRepository{}
	roomService:=room.NewRoomService(&roomRepo)
	rdb:= utils.CreateRedis()

	// var grpcClient pb.RemoteServerClient

	// var opts []grpc.DialOption
	// opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// grpcConn,err:= grpc.NewClient("localhost:50051",opts...)
	// if err!=nil {
	// 	log.Fatal(err)
	// }
	// grpcClient = pb.NewRemoteServerClient(grpcConn)
    // // return &client,conn
	
	
	

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    
    userId := r.URL.Query().Get("userId")
	roomId := r.URL.Query().Get("roomId")

		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Print(err)
			return
		}

		roomService.AddRoom(roomId)
		roomService.UpdateUser(userId,conn,roomId)


		defer func(){
            roomService.DeleteUser(userId,roomId)
			conn.Close()
			//grpcConn.Close()
		}()

			for {
	
				//publishing data
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					log.Print(err)
					break
				}
				sub:=pubsub.NewSubscribeServer(rdb,roomId,ctx)
                pubsub.PublishData(rdb,ctx,msg,roomId)
				//publishing data

				
				
				//subscribing data
			 subMessage,err := sub.RetrieveData(ctx)
             
			 if err!=nil {
				log.Print(err)
				return
			 }

			 fmt.Printf("Returned to client: %+v\n", subMessage)

			 for _,value:=range roomService.Room[subMessage.RoomId].Conn{

				marshaledData,err:=json.Marshal(subMessage)

				if err != nil {
					log.Print(err)
					break
				}

				err = wsutil.WriteServerMessage(value, op,marshaledData)
				
				if err != nil {
					log.Print(err)
					break
				}
				 fmt.Print(subMessage)
				// GRPCClient.SetPosition(grpcClient,&pb.PlayerPosition{
				// 	MetaId: subMessage.MetaId,
				// 	RoomId: subMessage.RoomId,
				// 	XPosition: subMessage.X_position,
				// 	YPosition: subMessage.Y_position,
				//  })
			 }
             	//subscribing data
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