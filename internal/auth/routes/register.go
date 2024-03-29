package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/auth/pb"
	"net/http"
)

type RegisterRequestBody struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c pb2.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb2.RegisterRequest{
		Name:     body.Name,
		Lastname: body.Lastname,
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
