package main

import (
	"context"
	echov1 "github.com/FlatMapIO/grpc-service-template/go/api/echo/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestInvokeVTGrpc(t *testing.T) {

	go StartGrpcServer()

	conn, err := grpc.DialContext(context.Background(), "localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	client := echov1.NewEchoServiceClient(conn)

	resp, err := client.Echo(context.Background(), &echov1.EchoRequest{
		Message: "Hello",
	})
	require.NoError(t, err)
	require.Equal(t, "Hello", resp.Message)
}