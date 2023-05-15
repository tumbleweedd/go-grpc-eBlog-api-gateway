package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/comment/pb"
	"net/http"
	"strconv"
)

func GetCommentsByPostId(ctx *gin.Context, client pb2.CommentServiceClient) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.GetCommentsByPostId(context.Background(), &pb2.GetCommentsByPostIdRequest{
		PostId: int64(postId),
	})

	ctx.JSON(int(res.Status), &res)
}
