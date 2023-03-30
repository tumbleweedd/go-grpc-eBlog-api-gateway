package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/comment/pb"
	"net/http"
	"strconv"
)

func GetCommentByCommentId(ctx *gin.Context, client pb.CommentServiceClient) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	commentId, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.GetCommentById(context.Background(), &pb.GetCommentByIdRequest{
		CommentId: int64(commentId),
		PostId:    int64(postId),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
