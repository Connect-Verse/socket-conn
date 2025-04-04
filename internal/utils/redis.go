package utils

import(
	"fmt"
	"github.com/redis/go-redis/v9"
)

func CreateRedis() *redis.Client {
	 client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", 
        DB:		  0,  
        Protocol: 2,  
    })

	fmt.Print("your redis client is successfully mounted")

	return client
}