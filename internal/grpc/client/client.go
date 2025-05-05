package GRPCClient

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/saransh-g1/socket-conn/internal/grpc"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr= flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func SetPosition(client pb.RemoteServerClient,position *pb.PlayerPosition){

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
    fmt.Print(position)
	
    set,err:= client.SetPositions(ctx,position)

	if err!=nil {
		log.Fatal("sorry error occurred while setting the position",err)
	}
    
	log.Println(set)

} 

type metaStruct struct{
	Id string
}

func FindPosition(client pb.RemoteServerClient, metaId string) (*pb.PositionResponse,error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

    result,err:= client.CheckPositions(ctx, &pb.MetaId{Id: metaId} )

	if err!=nil {
		log.Fatalf("sorry error occurred while setting the position")
		return nil,err
	}
	log.Println(result)

	return result,err
}


// func ClientServer() (*pb.RemoteServerClient,*grpc.ClientConn){	
	
// 	flag.Parse()
	
// }
