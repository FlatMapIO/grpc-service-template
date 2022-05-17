package main

import (
	"context"
	echov1 "github.com/FlatMapIO/grpc-service-template/go/api/echo/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	_ = cancel

	go StartGrpc(ctx)
	go StartGrpcGateway(ctx)
	select {}
}

func StartGrpc(ctx context.Context, ) {
	srv := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),

	)
	echov1.RegisterEchoServiceServer(srv, &echoServer{})
	reflection.Register(srv)
	const address = "0.0.0.0:8080"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	log.Info().Str("addr", address).Msg("Serving gRPC")
	if err := srv.Serve(lis); err != nil {
		log.Fatal().Err(err)
	}
}

func StartGrpcGateway(ctx context.Context) {
	const address = "0.0.0.0:8081"

	gwmux := runtime.NewServeMux()
	if err := echov1.RegisterEchoServiceHandlerServer(context.Background(), gwmux, &echoServer{}); err != nil {
		log.Fatal().Err(err).Msg("failed to register gRPC gateway")
	}
	srv := &http.Server{
		Addr:    address,
		Handler: gwmux,
	}
	log.Info().Str("addr", address).Msg("Serving gRPC Gateway")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal().Err(err)
	}
}