package main

import (
	"context"
	echov1 "github.com/FlatMapIO/grpc-service-template/go/api/echo/v1"
)

type echoService struct {
	echov1.UnimplementedEchoServiceServer
}

var _ echov1.EchoServiceServer = (*echoService)(nil)

func (e echoService) Echo(ctx context.Context, request *echov1.EchoRequest) (*echov1.EchoResponse, error) {
	return &echov1.EchoResponse{
		Message: request.Message,
	}, nil
}