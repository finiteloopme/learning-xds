package client

import (
	"context"
	"time"

	log "github.com/finiteloopme/goutils/pkg/log"

	grpcUtil "github.com/finiteloopme/goutils/pkg/grpc"
	"github.com/kelseyhightower/envconfig"

	userv1alpha1 "github.com/finiteloopme/xds-from-scratch/api/gen/proto/go/user/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunClient() {
	var config grpcUtil.GRPCConfig
	envconfig.Process("gcp", &config)
	conn, err := grpc.Dial(config.GRPC_Host+":"+config.GRPC_Port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	userClient := userv1alpha1.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	response, err := userClient.SayHello(ctx, &userv1alpha1.SayHelloRequest{User: &userv1alpha1.User{FirstName: "Kunal", LastName: "Limaye"}})
	if err != nil {
		log.Fatal(err)
	}
	log.Info(response.Msg)

}
