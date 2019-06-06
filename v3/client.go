package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	proto "gitee.com/go-microservice-study/v3/protoc"
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("hello.client")) // 客户端服务名称
	service.Init()
	helloservice := proto.NewHelloClient("hellooo", service.Client())
	// 请求服务
	req(helloservice, "world ^_^")
	// 以下是为了可以多次模拟请求写的, 可有可无
	for {
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("请输入: ")
		input.Scan()
		req(helloservice, input.Text())
	}
}

func req(helloservice proto.HelloClient, msg string) {
	res, err := helloservice.Ping(context.TODO(), &proto.Request{Name: msg})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Msg)
}
