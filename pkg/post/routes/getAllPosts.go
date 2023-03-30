package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/post/pb"
	"net/http"
)

func GetAllPosts(ctx *gin.Context, c pb.PostServiceClient) {
	res, err := c.GetAllPosts(context.Background(), &pb.GetAllPostsRequest{})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
