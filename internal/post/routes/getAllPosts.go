package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/post/pb"
	"net/http"
)

func GetAllPosts(ctx *gin.Context, c pb2.PostServiceClient) {
	res, err := c.GetAllPosts(context.Background(), &pb2.GetAllPostsRequest{})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
