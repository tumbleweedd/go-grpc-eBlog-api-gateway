package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/auth/routes"
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
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
