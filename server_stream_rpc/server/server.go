package main

import (
	"demo/MQ/model"
	pb "demo/server_stream_rpc/proto/helloworld"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

var Center *model.MQ

var count int

func (S *Server) SayHello(req *pb.HelloRequest, srv pb.Greeter_SayHelloServer) error {
	count++
	cli := model.NewClient(strconv.Itoa(count))
	Center.AddClient(cli)
	defer Center.UnSubscribe(cli.GetName())

	for data := range cli.Transport {
		fmt.Println(cli.GetName(), ":", data)
		if err := srv.Send(&pb.HelloReply{Message: data.(string)}); err != nil {
			return err
		}
	}
	return nil
}
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8503")
	if err != nil {
		log.Fatalf("failed listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	log.Printf("Server listen:%v", lis.Addr())

	Center = model.NewMQ()
	go func(mq *model.MQ) {
		for {
			time.Sleep(time.Second * 4)
			fmt.Println("____________________")
			mq.Notify("ss")
			fmt.Println("____________________")
		}
	}(Center)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed Server: %v", err)
	}
}
