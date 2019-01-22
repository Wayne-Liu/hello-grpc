package main

import (
	"google.golang.org/grpc"
	"time"
	"log"
	"github.com/Wayne-Liu/hello-grpc/model/gen"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

const (
	address     = "47.52.75.214:80"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := model.NewHelloServiceClient(conn)


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx,&empty.Empty{} )
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Response)



}
