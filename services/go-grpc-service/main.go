package main

import (
	"context"
	"log"
	"net"

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
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterGoGrpcServiceServer(grpcServer, &GoGrpcServiceServer{})
	grpcServer.Serve(lis)
}
