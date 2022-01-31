package main

import (
	"context"
	"log"
	chat "main/internal/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Not conn %v", err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{Body: "TEST"}

	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("Message error %v", err)
	}

	log.Printf("Response from server %s", response.Body)
}
