package server

import (
	"github-actions-nginx-demo/internal/service"
	"github-actions-nginx-demo/pb"
	"net"

	"google.golang.org/grpc"
)

func (s *server) NewTestGrpcServer() (func() error, *grpc.Server, error) {
	l, err := net.Listen("tcp", s.cfg.GRPC.Port)
	if err != nil {
		return nil, nil, err
	}

	server := grpc.NewServer()

	service := service.NewTestService(s.log, s.cfg)
	pb.RegisterTestServiceServer(server, service)

	go func() {
		s.log.Infof("Grpc server listening on port {%s}", s.cfg.GRPC.Port)
		server.Serve(l)
	}()

	s.log.Debugf("Хуита")

	return l.Close, server, nil
}
