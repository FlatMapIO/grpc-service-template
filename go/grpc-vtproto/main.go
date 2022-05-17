package main

import (
	echov1 "github.com/FlatMapIO/grpc-service-template/go/api/echo/v1"
	grpccodec "github.com/planetscale/vtprotobuf/codec/grpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	_ "google.golang.org/grpc/encoding/proto"
	"net"
)

func init() {
	log.Info().Msg("do package init")
	encoding.RegisterCodec(grpccodec.Codec{})
}

func main() {
	StartGrpcServer()
}

func StartGrpcServer() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	echov1.RegisterEchoServiceServer(srv, &echoService{})
	log.Info().Msg("Starting gRPC server")
	srv.Serve(lis)
}