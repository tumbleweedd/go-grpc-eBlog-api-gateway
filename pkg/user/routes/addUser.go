package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/user/pb"
	"net/http"
)

type Role string

type AddUserRequestBody struct {
	Name               string `json:"name" `
	Lastname           string `json:"lastname" `
	Username           string `json:"username" `
	Password           string `json:"password" `
	Email              string `json:"email" `
	Role               Role   `json:"role" `
	IsAccountNonLocked bool   `json:"is_account_non_locked" `
}

func AddUser(ctx *gin.Context, c pb.UserServiceClient) {
	body := AddUserRequestBody{}
	currentUserId, exist := ctx.Get("userId")
	if !exist {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.AddUser(context.Background(), &pb.AddUserRequest{
		CurrentUserId: currentUserId.(int64),
		Data: &pb.UserForAdminData{
			Name:               body.Name,
			Lastname:           body.Lastname,
			Username:           body.Username,
			Password:           body.Password,
			Email:              body.Email,
			Role:               string(body.Role),
			IsAccountNonLocked: body.IsAccountNonLocked,
		},
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
