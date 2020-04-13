package main

import (
	"context"
	"log"

	pb "github.com/drhelius/grpc-demo-user/server/grpc"
)

func main() {
	grpc.Serve()
}
