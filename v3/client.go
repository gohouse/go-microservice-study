package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "gohouse/go-microservice-study/v3/protoc"
	"os"
)


func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	req(greeter, "World")

	// 以下是为了可以多次模拟请求写的, 可有可无
	for {
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("请输入: ")
		input.Scan()
		req(greeter, input.Text())
	}
}

func req(helloservice proto.GreeterService, msg string) {
	res, err := helloservice.Hello(context.TODO(), &proto.HelloRequest{Name: msg})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Greeting)
}
