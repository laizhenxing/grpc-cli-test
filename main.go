package main

import (
	"context"
	"fmt"

	pb "greeter/pb"
	"github.com/micro/go-micro/v2"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9052", grpc.WithInsecure())
	if err != nil {
		fmt.Println("访问出错：", err)
	}
	defer conn.Close()

	msg := &pb.CommMessage{
		Uid: nil,
		All: true,
		Push: true,
		PushTitle: "test",
		PushBody: "This is a push test body",
		PushImage: "image",
	}
	svc := micro.NewService()
	svc.Init()
	c := pb.NewMessageService("dn.mc.srv.slope_msg", svc.Client())
	r, err := c.SendCommMessage(context.Background(), msg)
	if err != nil {
		fmt.Println("Call grpc-cli-test service fail", err)
	}

	fmt.Println("Get response: ", r)
}
