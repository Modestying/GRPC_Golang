package main

import (
	pb "demo/go/helloworld"
	"demo/internal/service"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":5602")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &service.HelloWorldService{})
	log.Printf("server listening at %v", lis.Addr())
	sysSignal := make(chan os.Signal)
	defer func(s chan os.Signal, server *grpc.Server) {
		close(s)
		log.Println("Close signal channel")
		server.Stop()
	}(sysSignal, s)
	var wait sync.WaitGroup
	wait.Add(1)
	SignalManage(&wait, sysSignal)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	wait.Wait()
}
func SignalManage(wait *sync.WaitGroup, s chan os.Signal) {
	//监听所有信号
	signal.Notify(s)
	//阻塞直到有信号传入
	log.Println("启动")

	go func(c chan os.Signal) {
		s := <-c
		log.Println("退出信号", s)
		wait.Done()
	}(s)
}
