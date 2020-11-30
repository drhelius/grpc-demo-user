package impl

import (
	"context"
	"log"
	"strconv"

	"github.com/Pallinder/go-randomdata"
	"github.com/drhelius/grpc-demo-proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	user.UnimplementedUserServiceServer
}

func (s *Server) Create(ctx context.Context, in *user.CreateUserReq) (*user.CreateUserResp, error) {

	log.Printf("[User] Create Req: %v", in.GetUser())

	r := &user.CreateUserResp{Id: strconv.Itoa(randomdata.Number(1000000))}

	err := failedContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("[User] Create Res: %v", r.GetId())

	return r, nil
}

func (s *Server) Read(ctx context.Context, in *user.ReadUserReq) (*user.ReadUserResp, error) {

	log.Printf("[User] Read Req: %v", in.GetId())

	r := &user.ReadUserResp{User: &user.User{Id: in.GetId(), Name: randomdata.FullName(randomdata.RandomGender), Email: randomdata.Email()}}

	err := failedContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("[User] Read Res: %v", r.GetUser())

	return r, nil
}

func failedContext(ctx context.Context) error {
	if ctx.Err() == context.Canceled {
		log.Printf("[Order] context canceled, stoping server side operation")
		return status.Error(codes.Canceled, "context canceled, stoping server side operation")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("[Order] dealine has exceeded, stoping server side operation")
		return status.Error(codes.DeadlineExceeded, "dealine has exceeded, stoping server side operation")
	}

	return nil
}
