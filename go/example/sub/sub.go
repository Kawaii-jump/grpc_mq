package main

import (
	"fmt"

	"github.com/Kawaii-jump/grpc_mq/go/client"
	"github.com/Kawaii-jump/grpc_mq/go/client/grpc"
)

func main() {
	GrpcSub()
}

func GrpcSub() {
	c := grpc.New(
		client.WithServers("http://10.224.205.72:8081"),
	)

	topic := "grpc"
	ch, err := c.Subscribe(topic)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}

	defer c.Unsubscribe(ch)

	i := 1
	for {
		select {
		case e := <-ch:
			fmt.Printf("sub topic:%s,\tnumber:%d,\tmessage:%s\n", topic, i, e)
			i++
		}
	}

}

func HttpSub() {
	c := client.New(
		client.WithServers("http://10.224.205.72:8081"),
	)

	topic := "http"

	ch, err := c.Subscribe(topic)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	defer c.Unsubscribe(ch)

	for {
		select {
		case e := <-ch:
			fmt.Printf("sub topic:%s,\tmessage:%s\n", topic, e)
		}
	}
}
