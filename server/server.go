package server

import (
	"context"
	"github-actions-nginx-demo/config"
	"github-actions-nginx-demo/logger"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	log logger.Logger
	cfg *config.Config
}

func NewServer(log logger.Logger, cfg *config.Config) *server {
	return &server{log: log, cfg: cfg}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	closeGrpcServer, grpcServer, err := s.NewTestGrpcServer()
	if err != nil {
		cancel()
		return err
	}

	defer closeGrpcServer()
	<-ctx.Done()

	grpcServer.GracefulStop()

	return nil
}
