package server

import (
	"context"
	"fmt"
	"io"

	grpcUtil "github.com/finiteloopme/goutils/pkg/grpc"
	log "github.com/finiteloopme/goutils/pkg/log"
	mathv1alpha1 "github.com/finiteloopme/xds-from-scratch/api/gen/proto/go/math/v1alpha1"
	"google.golang.org/grpc"
)

// todo: Fix hardcoding of Hello in this File
type MathService struct {
	// todo: fix hardcoding
	grpcUtil.UnimplementedGRPCServer
	mathv1alpha1.UnimplementedMathOperationServer
}

func (mathService MathService) Register(server *grpc.Server) {
	mathv1alpha1.RegisterMathOperationServer(server, &MathService{})
}

func (mathService *MathService) String() string {
	return "MathService"
}

func CalculatePowerOf(number, power float32) float32 {
	var product float32 = number
	if power > 1 {
		power -= 1
		product *= CalculatePowerOf(number, power)
	}

	return product
}

func (mathService *MathService) GetSquareOf(ctx context.Context, mathOp *mathv1alpha1.GetSquareOfRequest) (*mathv1alpha1.GetSquareOfResponse, error) {
	number := *&mathOp.Req.Content
	log.Info("Request to calculate square root of: " + fmt.Sprintf("%f", number))

	return &mathv1alpha1.GetSquareOfResponse{Res: &mathv1alpha1.Number{
		Content: CalculatePowerOf(number, 2),
	}}, nil
}

func (mathService *MathService) StreamSquareOf(stream mathv1alpha1.MathOperation_StreamSquareOfServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}
		number := in.GetReq().GetContent()
		stream.Send(&mathv1alpha1.StreamSquareOfResponse{
			Res: &mathv1alpha1.Number{
				Content: CalculatePowerOf(number, 2),
			},
		})

	}
}

func RunServer() error {
	var mathService MathService
	log.Info("Starting service: " + mathService.String())
	grpcUtil.RunGRPC(mathService)
	return nil
}
