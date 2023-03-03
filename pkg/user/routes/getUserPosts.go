package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/user/pb"
	"net/http"
)

func GetUserPosts(ctx *gin.Context, c pb.UserServiceClient) {
	username := ctx.Param("username")
	res, err := c.GetUserPosts(context.Background(), &pb.GetUserPostsRequest{
		Username: username,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
