package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/auth/pb"
	"net/http"
)

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb2.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb2.LoginRequest{
		Username: b.Username,
		Password: b.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
