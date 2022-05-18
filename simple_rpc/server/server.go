package main

import (
	"context"
	pb "demo/simple_rpc/proto/helloworld"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (S *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println(md.Get("x-real-ip"))
		fmt.Println(md.Get("x-forwarded-for"))
	}
	return &pb.HelloReply{Message: "Hello," + in.GetName()}, nil
}
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8503")
	if err != nil {
		log.Fatalf("failed listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	log.Printf("Server listen:%v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed Server: %v", err)
	}
}
