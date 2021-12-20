package main

import (
	pbv1 "github.com/danlmarmot/grpc-v-greeter/protogen/v1/greeter"
	pbv2 "github.com/danlmarmot/grpc-v-greeter/protogen/v2/greeter"
	"log"
	"net"

	greeterv1 "github.com/danlmarmot/grpc-v-greeter/pkg/service/greeter/v1"
	greeterv2 "github.com/danlmarmot/grpc-v-greeter/pkg/service/greeter/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("unable to listen on port ", port)
	}
	s := grpc.NewServer()
	reflection.Register(s)

	// Register API v1
	greeterV1, err := greeterv1.NewService()
	if err != nil {
		log.Fatal("unable to initialise v1 service")
	}
	pbv1.RegisterGreeterServer(s, greeterV1)

	// Register API v2
	greeterV2, err := greeterv2.NewService()
	if err != nil {
		log.Fatal("unable to initialise v2 service")
	}
	pbv2.RegisterGreeterServer(s, greeterV2)

	log.Printf("listening on port %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
