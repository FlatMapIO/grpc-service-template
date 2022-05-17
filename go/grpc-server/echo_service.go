package main

import (
	"context"
	"fmt"
	echov1 "github.com/FlatMapIO/grpc-service-template/go/api/echo/v1"
	"github.com/rs/zerolog/log"
)

type echoServer struct {
	echov1.UnimplementedEchoServiceServer
}

func (e echoServer) Echo(ctx context.Context, request *echov1.EchoRequest) (*echov1.EchoResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	return &echov1.EchoResponse{
		Message: request.Message,
	}, nil
}

func (e echoServer) EchoStreaming(request *echov1.ServerStreamingEchoRequest, server echov1.EchoService_EchoStreamingServer) error {
	count := 0
	for {
		count++
		if count > 10 {
			return nil
		}

		if err := server.Send(&echov1.ServerStreamingEchoResponse{
			Message: fmt.Sprintf("%d: %s", count, request.Message),
		}); err != nil {
			log.Error().Err(err).Msg("failed to send streaming response")

			return err
		}
	}
}

var _ echov1.EchoServiceServer = (*echoServer)(nil)