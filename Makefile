BINARY_PATH=bin

MAIN_PATH=cmd/grpc-demo-user/main.go
BINARY_NAME=$(BINARY_PATH)/grpc-demo-user

PROTO_FILES_PATH=../grpc-demo-proto
PROTO_OUT_PATH=internal/grpc

build:
	go build -o $(BINARY_NAME) -race $(MAIN_PATH)

clean:
	go clean $(MAIN_PATH)
	rm -f $(BINARY_PATH)/*

build-proto:
	mkdir -p $(PROTO_OUT_PATH)
	protoc -I $(PROTO_FILES_PATH) -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:$(PROTO_OUT_PATH) user/user.proto

clean-proto:
	rm -f $(PROTO_OUT_PATH)/*.pb.go

# Proto parameters
# PROTO_FILES_PATH=proto
# PROTO_OUT=genproto/go

# gen-proto:
# 	protoc -I $(PROTO_FILES_PATH) -I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:$(PROTO_OUT) --grpc-gateway_out=logtostderr=true:$(PROTO_OUT) $(PROTO_FILES_PATH)/*.proto

# clean-proto:
# 	rm -f $(PROTO_OUT)/*.pb.go