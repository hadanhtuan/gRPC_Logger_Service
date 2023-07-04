package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "logger/protobuf"
	"time"
)

func main2() {

	conn, err := grpc.Dial("localhost:5151", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()

	c := pb.NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result, err2 := c.WriteLog(ctx, &pb.LogRequest{
		LogEntry: &pb.LogEntry{
			Name: "Client test",
			Data: "motherfucker",
		},
	})

	if err2 != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("Done, api called motherfucker")
		fmt.Println(result.Message)
	}

}
