package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-demo/generated"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := generated.NewRepositoryServiceClient(conn)

	response, err := client.Receive(context.Background(), &generated.Mail{Text: "Send To gRPC Server"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Text)

}
