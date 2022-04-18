package main

import (
	"flag"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("Read from tcp server failed,err:", err)
			break
		}
		data := string(buf[:n])
		fmt.Printf("Recived from client,data:%s\n", data)
	}
}

var port int
func init() {
	flag.IntVar(&port, "port", 1234, "Just for demo")
}

// ./main.exe 1 2 "third" --port=15
func main() {
	flag.Parse()
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("port=", port)
	//// 监听TCP 服务端口
	//listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))
	//if err != nil {
	//	fmt.Println("Listen tcp server failed,err:", err)
	//	return
	//}
	//
	//for {
	//	// 建立socket连接
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		fmt.Println("Listen.Accept failed,err:", err)
	//		continue
	//	}
	//
	//	// 业务处理逻辑
	//	go process(conn)
	//}
}
