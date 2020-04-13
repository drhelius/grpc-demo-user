package impl

import (
	"context"
	"log"

	pb "github.com/drhelius/grpc-demo-user/internal/grpc/user"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Create(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserResp, error) {

	log.Printf("Received: %v", in.GetUser())

	return &pb.CreateUserResp{Id: "testid"}, nil
}

func (s *server) Read(ctx context.Context, in *pb.ReadUserReq) (*pb.ReadUserResp, error) {

	log.Printf("Received: %v", in.GetId())

	return &pb.ReadUserResp{User: &pb.User{Id: "demoid", Name: "demoname", Email: "demoemail"}}, nil
}
