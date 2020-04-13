package impl

import (
	"context"
	"log"

	pb "github.com/drhelius/grpc-demo-user/internal/grpc/user"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s *Server) Create(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserResp, error) {

	log.Printf("Received: %s", in.GetUser())

	return &pb.CreateUserResp{Id: "testid"}, nil
}

func (s *Server) Read(ctx context.Context, in *pb.ReadUserReq) (*pb.ReadUserResp, error) {

	log.Printf("Received: %v", in.GetId())

	return &pb.ReadUserResp{User: &pb.User{Id: "demoid", Name: "demoname", Email: "demoemail"}}, nil
}
