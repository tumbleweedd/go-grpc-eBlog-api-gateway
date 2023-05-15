package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/user/pb"
	"net/http"
)

func GetUserProfile(ctx *gin.Context, c pb2.UserServiceClient) {
	username := ctx.Param("username")

	res, err := c.GetUserProfile(context.Background(), &pb2.GetUserProfileRequest{
		Username: username,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
