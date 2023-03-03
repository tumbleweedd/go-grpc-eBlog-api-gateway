package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/user/pb"
	"net/http"
)

func GetLoggedUserProfile(ctx *gin.Context, c pb.UserServiceClient) {
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := c.GetLoggedUserProfile(context.Background(), &pb.GetLoggedUserProfileRequest{
		UserId: userId.(int64),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
