package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/user/pb"
	"net/http"
)

func DeleteUserProfile(ctx *gin.Context, c pb.UserServiceClient) {
	username := ctx.Param("username")
	currentUserId, exist := ctx.Get("userId")
	if !exist {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := c.DeleteUser(context.Background(), &pb.DeleteUserRequest{
		Username: username,
		UserId:   currentUserId.(int64),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	}

	ctx.JSON(int(res.Status), &res)
}
