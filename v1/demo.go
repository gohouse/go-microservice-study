package main

import (
	"fmt"
	"gohouse/go-microservice-study/v1/protoc"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	var mobile = &protoc.Mobile{Brand: "xiaomi"}

	// 转码
	data, err := proto.Marshal(mobile)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// 解码
	var newMobile protoc.Mobile
	err = proto.Unmarshal(data, &newMobile)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// 判断转码前后转码后解码的model是否一致
	if mobile.GetBrand() != newMobile.GetBrand() {
		log.Fatalf("data mismatch %q != %q", mobile.GetBrand(), newMobile.GetBrand())
	}

	fmt.Printf("原始brand: %s \n转码后解码的brand: %s\n", mobile.GetBrand(), newMobile.GetBrand())
}
