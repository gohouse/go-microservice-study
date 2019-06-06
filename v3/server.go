package main

import (
	"context"
	"fmt"

	proto "gitee.com/go-microservice-study/v3/protoc"

	micro "github.com/micro/go-micro"
)

type Hello struct{}

func (h *Hello) Ping(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.Msg = "Hello " + req.Name
	return nil
}
func main() {
	service := micro.NewService(
		micro.Name("hellooo"), // 服务名称
	)
	service.Init()
	proto.RegisterHelloHandler(service.Server(), new(Hello))
	err := service.Run()
	if err != nil {
		fmt.Println(err)
	}
}
