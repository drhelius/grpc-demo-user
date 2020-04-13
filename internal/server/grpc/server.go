package grpc

import (
	"log"
	"net"

	pb "github.com/drhelius/grpc-demo-user/internal/grpc/user"
	"github.com/drhelius/grpc-demo-user/internal/impl/user"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func Serve() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
