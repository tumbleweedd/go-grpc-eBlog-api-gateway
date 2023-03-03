package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/user/pb"
	"net/http"
)

func GetUserList(ctx *gin.Context, c pb.UserServiceClient) {
	res, err := c.GetUserList(context.Background(), &pb.GetUserListRequest{})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
