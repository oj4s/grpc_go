package main

import (
	"context"
	"log"

	"github.com/oj4s/protobuf/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", errr)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	message := chat.Message {
		Body: "Hello from the client",
	}
	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error while calling SayHello: %s",err)
	}

	log.Printf("Response from server: %s",response.Body)
