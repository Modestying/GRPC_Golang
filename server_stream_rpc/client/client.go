package main

import (
	"context"
	pb "demo/server_stream_rpc/proto/helloworld"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "0.0.0.0:8503", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		if res, err := r.Recv(); err == nil {
			log.Println(res)
		} else if err == io.EOF {
			break
		} else {
			log.Println("err: ", err.Error())
		}

	}
}
