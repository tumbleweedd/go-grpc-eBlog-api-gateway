package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/user/pb"
	"net/http"
)

type UserUpdateRequestBody struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func UpdateUserProfile(ctx *gin.Context, c pb2.UserServiceClient) {
	body := UserUpdateRequestBody{}

	currentUserId, exist := ctx.Get("userId")
	if !exist {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	username := ctx.Param("username")

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UpdateUser(context.Background(), &pb2.UpdateUserRequest{
		UserId:   currentUserId.(int64),
		Username: username,
		Data: &pb2.UserUpdateData{
			Username: body.Username,
			Email:    body.Email,
			Password: body.Password,
		},
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}
