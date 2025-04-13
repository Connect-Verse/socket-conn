package pubsub

import (
	"context"
	"encoding/json"
	"time"
	"fmt"
	"log"
	"github.com/redis/go-redis/v9"
	"github.com/saransh-g1/socket-conn/internal/data"
)


type Subsribe interface{
	RetrieveData(context context.Context) (data.PositionResponse,error)
}


type SubscribeService struct{
	Sub *redis.PubSub
	Positions chan data.PositionResponse
}

func NewSubscribeServer(rdb *redis.Client,roomId string,context context.Context) SubscribeService{
	sub := rdb.Subscribe(context,roomId)

	positions:= make(chan data.PositionResponse,100)
	service :=SubscribeService{Sub:sub, Positions: positions}

	service.listen()

  return service
}


func (s *SubscribeService) listen() {
	ch := s.Sub.Channel()

	go func(){
		for msg := range ch {
			var pos data.PositionResponse
			err := json.Unmarshal([]byte(msg.Payload), &pos)
			if err != nil {
				log.Println("Unmarshal error:", err)
				continue
			}
			//fmt.Printf("Received: %+v\n", pos)
			s.Positions<-pos
		}
	}() 
}


func (s *SubscribeService) RetrieveData(context context.Context) (data.PositionResponse,error){
	select {
	case pos := <-s.Positions:
		fmt.Printf("Returned to client: %+v\n", pos)
		return pos, nil
	case <-time.After(5 * time.Second): // timeout if no data
		return data.PositionResponse{}, fmt.Errorf("timeout: no data received")
	}
		}


