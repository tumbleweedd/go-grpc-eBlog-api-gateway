package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/comment/pb"
	"net/http"
	"strconv"
)

type AddCommentRequestBody struct {
	Body string `json:"body"`
}

func AddComment(ctx *gin.Context, client pb.CommentServiceClient) {
	request := AddCommentRequestBody{}

	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := client.AddComment(context.Background(), &pb.AddCommentRequest{
		UserId: userId.(int64),
		PostId: int64(postId),
		Body:   request.Body,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
