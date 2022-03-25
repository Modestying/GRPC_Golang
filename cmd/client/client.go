package main

import (
	"context"
	pb "demo/go/helloworld"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":5602", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	s := make(chan os.Signal)

	defer func(conn *grpc.ClientConn, s chan os.Signal) {
		conn.Close()
		log.Println("Grpc connection closed")
		close(s)
		log.Println("close channel")
	}(conn, s)

	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var wait sync.WaitGroup
	wait.Add(1)

	go SignalManage(&wait, s)

	name := "jack"
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
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
