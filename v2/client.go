package main

import (
	"bufio"
	"context"
	"fmt"
	pb "gohouse/go-microservice-study/v2/protoc"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	address = "localhost:50051"
)

func main() {
	//建立链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	//// 10秒的上下文
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	//defer cancel()
	// 请求服务
	req(c, "world ^_^")
	// 以下是为了可以多次模拟请求写的, 可有可无
	for {
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("请输入: ")
		input.Scan()
		req(c, input.Text())
	}
}

func req(c pb.GreeterClient, name string) {
	r, err := c.SayHello(context.TODO(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
