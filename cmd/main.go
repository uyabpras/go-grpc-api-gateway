package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/auth"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/config"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/order"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/product"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("failed to load config: %v\n", err)
	}

	r := gin.Default()

	authsvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authsvc)
	order.RegisterRoutes(r, &c, &authsvc)

	r.Run(c.Port)
}
