package service

import (
	"context"
	"fmt"
	"github-actions-nginx-demo/config"
	"github-actions-nginx-demo/logger"
	"github-actions-nginx-demo/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TestService struct {
	pb.UnimplementedTestServiceServer

	log logger.Logger
	cfg *config.Config
}

func NewTestService(log logger.Logger, cfg *config.Config) *TestService {
	return &TestService{log: log, cfg: cfg}
}

func (t *TestService) Test(ctx context.Context, req *pb.TestReq) (*pb.TestRes, error) {
	testText := req.GetTestText()
	if testText == "error" || testText == "" {
		return nil, status.Error(codes.InvalidArgument, "testText is empty or equal error")
	}

	return &pb.TestRes{
		TestText: fmt.Sprintf("This is your testText {%s} alright?", testText),
	}, nil
}
