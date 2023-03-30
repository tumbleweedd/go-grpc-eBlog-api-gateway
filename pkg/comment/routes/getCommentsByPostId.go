package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/comment/pb"
	"net/http"
	"strconv"
)

func GetCommentsByPostId(ctx *gin.Context, client pb.CommentServiceClient) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.GetCommentsByPostId(context.Background(), &pb.GetCommentsByPostIdRequest{
		PostId: int64(postId),
	})

	ctx.JSON(int(res.Status), &res)
}
