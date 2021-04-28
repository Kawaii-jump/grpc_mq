package grpc

import (
	"net"

	"github.com/Kawaii-jump/grpc_mq/proto/grpc/mq"
	"github.com/Kawaii-jump/grpc_mq/server"
	"github.com/Kawaii-jump/grpc_mq/server/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type grpcServer struct {
	options *server.Options
}

func (g *grpcServer) Run() error {
	l, err := net.Listen("tcp", g.options.Address)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption

	// tls enabled
	if g.options.TLS != nil {
		creds, err := credentials.NewServerTLSFromFile(
			g.options.TLS.CertFile,
			g.options.TLS.KeyFile,
		)
		if err != nil {
			return err
		}
		opts = append(opts, grpc.Creds(creds))
	} else {
		// generate tls config
		addr, err := util.Address(g.options.Address)
		if err != nil {
			return err
		}

		cert, err := util.Certificate(addr)
		if err != nil {
			return err
		}

		creds := credentials.NewServerTLSFromCert(&cert)
		opts = append(opts, grpc.Creds(creds))
	}

	// new grpc server
	srv := grpc.NewServer(opts...)

	// register MQ server
	mq.RegisterMQServer(srv, new(handler))

	// serve
	return srv.Serve(l)
}

func New(opts ...server.Option) *grpcServer {
	options := new(server.Options)
	for _, o := range opts {
		o(options)
	}
	return &grpcServer{
		options: options,
	}
}
