package main

import (
	"fmt"
	"time"

	"github.com/Kawaii-jump/grpc_mq/go/client"
	"github.com/Kawaii-jump/grpc_mq/go/client/grpc"
)

func main() {
	GrpcPub()
}

func GrpcPub() {
	c := grpc.New(
		client.WithServers("http://10.224.205.72:8081"),
	)
	tick := time.NewTicker(time.Second)

	topic := "grpc"
	message := []byte(`grpc message`)

	i := 1

	for _ = range tick.C {
		if err := c.Publish("grpc", message); err == nil {
			fmt.Printf("pub topic:%s,\tnumber:%d,\tmessage:%s\n", topic, i, message)
			i++
		} else {
			fmt.Printf("error:%v\n", err)
			break
		}
	}

}

func HttpPub() {
	c := client.New(
		client.WithServers("http://10.224.205.72:8081"),
	)
	tick := time.NewTicker(time.Second)

	topic := "http"
	message := []byte(`http message`)

	for _ = range tick.C {
		if err := c.Publish(topic, message); err == nil {
			fmt.Printf("pub topic:%s,\tmessage:%s\n", topic, message)
		} else {
			fmt.Printf("error:%v\n", err)
			break
		}
	}
}
