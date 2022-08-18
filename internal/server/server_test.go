package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	log "github.com/finiteloopme/goutils/pkg/log"
	mathv1alpha1 "github.com/finiteloopme/xds-from-scratch/api/gen/proto/go/math/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

func dialer() func(context.Context, string) (net.Conn, error) {
	// RunServer()
	listener := bufconn.Listen(bufSize)
	var mathService MathService
	log.Info("Starting service: " + mathService.String())
	server := grpc.NewServer()

	mathService.Register(server)
	// mathv1alpha1.RegisterMathOperationServer(server, &MathService{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestPowerOf(t *testing.T) {
	number1 := 2.0
	number2 := 11.0
	expected1 := 4.0
	expected2 := 14641.0

	rec1 := CalculatePowerOf(float32(number1), 2)

	if rec1 != float32(expected1) {
		t.Fatalf("Received %v, expected %v", rec1, expected1)
	}

	rec2 := CalculatePowerOf(float32(number2), 4)
	if rec2 != float32(expected2) {
		t.Fatalf("Received %v, expected %v", rec2, expected2)
	}

}

func TestSquareOf(t *testing.T) {

	tests := []struct {
		name string
		req  mathv1alpha1.GetSquareOfRequest
		res  mathv1alpha1.GetSquareOfResponse
	}{
		{
			"Square of 5",
			mathv1alpha1.GetSquareOfRequest{Req: &mathv1alpha1.Number{Content: 5}},
			mathv1alpha1.GetSquareOfResponse{Res: &mathv1alpha1.Number{Content: 25}},
		},
		{
			"Square of 7",
			mathv1alpha1.GetSquareOfRequest{Req: &mathv1alpha1.Number{Content: 7}},
			mathv1alpha1.GetSquareOfResponse{Res: &mathv1alpha1.Number{Content: 49}},
		},
		{
			"Square of 11",
			mathv1alpha1.GetSquareOfRequest{Req: &mathv1alpha1.Number{Content: 11}},
			mathv1alpha1.GetSquareOfResponse{Res: &mathv1alpha1.Number{Content: 121}},
		},
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := mathv1alpha1.NewMathOperationClient(conn)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := client.GetSquareOf(ctx, &tt.req)
			if res.GetRes().GetContent() != tt.res.GetRes().GetContent() {
				t.Fatalf("Expected %v, received %v", tt.res.GetRes().GetContent(), res.GetRes().GetContent())
			}
		})
	}

}

func TestStreamSquareOf(t *testing.T) {
	type testType struct {
		name string
		req  mathv1alpha1.GetSquareOfRequest
		res  mathv1alpha1.GetSquareOfResponse
	}
	// var idx int = 100
	var tests [5000]testType
	// for index, test := range tests {
	for index := 0; index < 5000; index++ {
		tests[index].name = "Square of: " + fmt.Sprint(index)
		tests[index].req = mathv1alpha1.GetSquareOfRequest{Req: &mathv1alpha1.Number{Content: float32(index)}}
		tests[index].res = mathv1alpha1.GetSquareOfResponse{Res: &mathv1alpha1.Number{Content: float32(index * index)}}
	}
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	clientCtx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	client := mathv1alpha1.NewMathOperationClient(conn)
	stream, err := mathv1alpha1.MathOperationClient.StreamSquareOf(client, clientCtx)
	if err != nil {
		t.Errorf("Error creating gRPC streaming client: %v", err)
	}
	// client := mathv1alpha1.NewMathOperationClient(conn)
	waitResponse := make(chan error)
	// go routine to receive responses
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Info("No more responses")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("Cannot receive stream from server")
			}
			log.Info("Received response: " + fmt.Sprint(res.GetRes().GetContent()))
		}
	}()

	// go routine to send requests
	go func() {
		for _, test := range tests {
			log.Info("Sending square request for: " + fmt.Sprint(test.req.GetReq().GetContent()))
			stream.Send((*mathv1alpha1.StreamSquareOfRequest)(&test.req))
		}
		err = stream.CloseSend()
		if err != nil {
			waitResponse <- fmt.Errorf("cannot close send: %v", err)
		}
	}()

	err = <-waitResponse
	if err != nil {
		log.Fatal(err)
	}
}
