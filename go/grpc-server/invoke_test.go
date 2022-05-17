package main

import (
	"bytes"
	"context"
	echov1 "github.com/FlatMapIO/grpc-service-template/go/api/echo/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func connect(t *testing.T) grpc.ClientConnInterface {

	conn, err := grpc.DialContext(context.Background(), "localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return conn
}

func TestInvokeGrpcEcho(t *testing.T) {
	conn := connect(t)
	echo := echov1.NewEchoServiceClient(conn)

	resp, err := echo.Echo(context.Background(), &echov1.EchoRequest{Message: "Hello, world!"})
	require.NoError(t, err)
	require.Equal(t, "Hello, world!", resp.Message)
}

func TestInvokeGrpcGatewayEcho(t *testing.T) {

	const message = `{"message":"Hello, world!"}`
	req := http.Request{
		Method: "POST",
		URL: &url.URL{
			Scheme: "http",
			Host:   "localhost:8081",
			Path:   "/v1/example/echo",
		},
		Body: io.NopCloser(bytes.NewBuffer([]byte(message))),
	}
	resp, err := http.DefaultClient.Do(&req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	raw, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	t.Logf("server reply: %s", string(raw))
	require.Equal(t, message, string(raw))
}

func TestInvokeStream(t *testing.T) {
	conn := connect(t)
	echo := echov1.NewEchoServiceClient(conn)
	h, err := echo.EchoStreaming(context.Background(), &echov1.ServerStreamingEchoRequest{Message: "Hello, world!"})
	require.NoError(t, err)
	for {
		resp, err := h.Recv()
		if err == io.EOF {
			t.Log("end of stream")
			break
		}
		require.NoError(t, err)
		require.True(t, strings.HasSuffix(resp.Message, "Hello, world!"))
		t.Log(resp.Message)
	}
}