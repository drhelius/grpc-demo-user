package impl

import (
	"context"
	"log"

	"github.com/drhelius/grpc-demo-proto/user"
)

type Server struct {
	user.UnimplementedUserServiceServer
}

func (s *Server) Create(ctx context.Context, in *user.CreateUserReq) (*user.CreateUserResp, error) {

	log.Printf("Received: %s", in.GetUser())

	return &user.CreateUserResp{Id: "testid"}, nil
}

func (s *Server) Read(ctx context.Context, in *user.ReadUserReq) (*user.ReadUserResp, error) {

	log.Printf("Received: %v", in.GetId())

	return &user.ReadUserResp{User: &user.User{Id: "demoid", Name: "demoname", Email: "demoemail"}}, nil
}
