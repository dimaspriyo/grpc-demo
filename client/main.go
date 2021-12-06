package main

import (
	"github.com/go-chi/chi/v5"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"grpc-demo/generated"
	"log"
	"net/http"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("dns:///server:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := generated.NewRepositoryServiceClient(conn)

	r := chi.NewRouter()
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {

		response, err := client.Receive(context.Background(), &generated.Mail{Text: "Send To gRPC Server"})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Response from server: %s", response.Text)
	})
	http.ListenAndServe(":3000", r)

}
