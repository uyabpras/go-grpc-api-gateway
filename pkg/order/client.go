package order

import (
	"fmt"

	"github.com/uyabpras/go-grpc-api-gateway/pkg/config"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct{
	Client pb.OrderServiceClient
}

func InitServiceClient (c *config.Config) pb.OrderServiceClient{
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("couldn't connect to order service: ", err)
	}

	return pb.NewOrderServiceClient(cc)
}