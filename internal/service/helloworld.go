package service

import (
	"context"
	pb "demo/go/helloworld"
)

type HelloWorldService struct {
	pb.UnimplementedGreeterServer
}

func NewHelloWorldServer() *HelloWorldService {
	return &HelloWorldService{}
}

func (S *HelloWorldService) SayHello(context context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello" + req.GetName()}, nil
}
