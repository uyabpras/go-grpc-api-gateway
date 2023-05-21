package auth

import (
	"fmt"

	"github.com/uyabpras/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct{
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient{
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("couldn't not connect to auth service: ", err)
	}

	return pb.NewAuthServiceClient(cc)
}