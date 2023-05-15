package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/user/pb"
	"net/http"
)

func GetUserList(ctx *gin.Context, c pb2.UserServiceClient) {
	res, err := c.GetUserList(context.Background(), &pb2.GetUserListRequest{})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
