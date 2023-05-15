package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/post/pb"
	"net/http"
	"strconv"
)

func DeletePostById(ctx *gin.Context, client pb2.PostServiceClient) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.DeletePost(context.Background(), &pb2.DeletePostRequest{
		PostId: int64(postId),
	})

	ctx.JSON(int(res.Status), &res)
}
