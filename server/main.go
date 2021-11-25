package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/generated"
	"log"
	"net"
)

func main() {

	fmt.Println("gRPC Start !")

	mailImpl := generated.MailImpl{}
	grpcServer := grpc.NewServer()
	generated.RegisterRepositoryServiceServer(grpcServer, mailImpl)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
