package http

import "github.com/Kawaii-jump/grpc_mq/go/client"

// New returns a http client
func New(opts ...client.Option) client.Client {
	return client.New(opts...)
}
