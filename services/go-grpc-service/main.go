package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	proto "github.com/tomhaddad/grafana-beyla-kubernetes-test/proto/gen/go/go/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GoGrpcServiceServer struct {
	proto.UnimplementedGoGrpcServiceServer
}

func (s *GoGrpcServiceServer) Forward(ctx context.Context, request *proto.ForwardRequest) (*proto.ForwardResponse, error) {
	conn, err := grpc.Dial("go-grpc-service-2:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	client := proto.NewGoGrpcServiceClient(conn)

	client.Receive(context.Background(), &proto.ReceiveRequest{
		Message: request.Message,
	})

	return &proto.ForwardResponse{
		Message: fmt.Sprintf("Forwarded message: %s", request.Message),
	}, nil
}

func (s *GoGrpcServiceServer) Receive(ctx context.Context, request *proto.ReceiveRequest) (*proto.ReceiveResponse, error) {
	return &proto.ReceiveResponse{
		Message: fmt.Sprintf("Received message: %s", request.Message),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterGoGrpcServiceServer(grpcServer, &GoGrpcServiceServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
