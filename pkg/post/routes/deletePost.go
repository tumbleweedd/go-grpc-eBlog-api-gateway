package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/post/pb"
	"net/http"
	"strconv"
)

func DeletePostById(ctx *gin.Context, client pb.PostServiceClient) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.DeletePost(context.Background(), &pb.DeletePostRequest{
		PostId: int64(postId),
	})

	ctx.JSON(int(res.Status), &res)
}
