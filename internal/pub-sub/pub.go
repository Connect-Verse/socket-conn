package pubsub

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func PublishData(rdb *redis.Client,context context.Context,message []byte,roomId string) (error){

    err := rdb.Publish(context,roomId, message).Err()
	if err!=nil{
		fmt.Print("error occured while publishing the message")
		return err
	}
	return nil
}