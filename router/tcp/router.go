package router

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func Start() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	SimpleInit().simpleRouters(s)

	log.Println("gRPC server is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
