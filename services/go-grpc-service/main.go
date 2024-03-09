package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	proto "github.com/tomhaddad/grafana-beyla-kubernetes-test/proto/gen/go/go/v1"
	"google.golang.org/grpc"
)

type GoGrpcServiceServer struct {
	proto.UnimplementedGoGrpcServiceServer
}

func (s *GoGrpcServiceServer) Forward(ctx context.Context, request *proto.ForwardRequest) (*proto.ForwardResponse, error) {
	return &proto.ForwardResponse{}, nil
}

func (s *GoGrpcServiceServer) Receive(ctx context.Context, request *proto.ReceiveRequest) (*proto.ReceiveResponse, error) {
	return &proto.ReceiveResponse{}, nil
}

func NewGoGrpcServiceServer() GoGrpcServiceServer {
	return GoGrpcServiceServer{}
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	fmt.Printf("Serving on port %s\n", port)
	startServer(port)
}

func startServer(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterGoGrpcServiceServer(grpcServer, &GoGrpcServiceServer{})
	grpcServer.Serve(lis)
}
