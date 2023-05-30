package order

import (
	"github.com/gin-gonic/gin"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/auth"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/config"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/order/routes"
)

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}

func RegisterRoutes(r *gin.Engine, c *config.Config, authsvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authsvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.Authrequired)
	routes.POST("/", svc.CreateOrder)
}
