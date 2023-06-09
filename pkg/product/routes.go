package product

import (
	"github.com/gin-gonic/gin"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/auth"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/config"
	"github.com/uyabpras/go-grpc-api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authsvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authsvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.Authrequired)
	routes.POST("/", svc.CreateProduct)
	routes.GET("/find/:id", svc.FindOne)
	routes.GET("/all", svc.ListProduk)
	routes.GET("/download", svc.DownloadProduct)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) ListProduk(ctx *gin.Context) {
	routes.ListProduct(ctx, svc.Client)
}

func (svc *ServiceClient) DownloadProduct(ctx *gin.Context) {
	routes.DownloadProduct(ctx, svc.Client)
}
