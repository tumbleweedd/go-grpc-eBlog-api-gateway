package auth

import (
	"github.com/gin-gonic/gin"
	routes2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/auth/routes"
)

func RegisterRoutes(r *gin.Engine, authUrl string) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(authUrl),
	}

	routes := r.Group("/auth")
	{
		routes.POST("/sign-up", svc.Register)
		routes.POST("/sign-in", svc.Login)
	}

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes2.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes2.Login(ctx, svc.Client)
}
