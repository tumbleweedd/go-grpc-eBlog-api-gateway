package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/user/pb"
	"net/http"
)

func GetLoggedUserProfile(ctx *gin.Context, c pb2.UserServiceClient) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := c.GetLoggedUserProfile(context.Background(), &pb2.GetLoggedUserProfileRequest{
		UserId: userId.(int64),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
