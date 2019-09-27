package main

import (
	"gitee.com/tmp/openbilibili/app/admin/main/apm/service"
	proto "gohouse/go-microservice-study/v3/protoc"
	micro "github.com/micro/go-micro"
)

func main() {
	p := micro.NewPublisher("events", service.Client())
	p.Publish(context.TODO(), &proto.Event{Name: "event"})
}
