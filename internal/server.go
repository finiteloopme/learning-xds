package internal

import (
	"context"

	grpcUtil "github.com/finiteloopme/goutils/pkg/grpc"
	log "github.com/finiteloopme/goutils/pkg/log"
	userv1alpha1 "github.com/finiteloopme/xds-from-scratch/api/gen/proto/go/user/v1alpha1"
	"google.golang.org/grpc"
)

// todo: Fix hardcoding of Hello in this File
type MyHelloService struct {
	// todo: fix hardcoding
	userv1alpha1.UnimplementedHelloServiceServer
	grpcUtil.UnimplementedGRPCServer
}

func (quoteService MyHelloService) Register(server *grpc.Server) {
	userv1alpha1.RegisterHelloServiceServer(server, &MyHelloService{})
}

func (helloService *MyHelloService) SayHello(ctx context.Context, user *userv1alpha1.SayHelloRequest) (*userv1alpha1.SayHelloResponse, error) {
	name := user.User.FirstName + " " + user.User.LastName
	log.Info("name: " + name)
	if name == " " {
		name = "World"
	}
	return &userv1alpha1.SayHelloResponse{
		Msg: "Hello " + name,
	}, nil
}

func RunServer() error {
	var userService MyHelloService
	grpcUtil.RunGRPC(userService)
	return nil
}
