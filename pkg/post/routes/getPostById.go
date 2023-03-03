package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/post/pb"
	"net/http"
	"strconv"
)

func GetPostById(ctx *gin.Context, c pb.PostServiceClient) {
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.GetPostById(context.Background(), &pb.GetPostByIdRequest{
		PostId: int64(postId),
	})

	ctx.JSON(int(res.Status), &res)

}
