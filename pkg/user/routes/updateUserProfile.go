package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/user/pb"
	"net/http"
)

type UserUpdateRequestBody struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func UpdateUserProfile(ctx *gin.Context, c pb.UserServiceClient) {
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

	res, err := c.UpdateUser(context.Background(), &pb.UpdateUserRequest{
		UserId:   currentUserId.(int64),
		Username: username,
		Data: &pb.UserUpdateData{
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
