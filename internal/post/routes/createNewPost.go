package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	pb2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/post/pb"
	"net/http"
)

type CreatPostRequestBody struct {
	Body     string   `json:"body"`
	Head     string   `json:"title"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

func CreateNewPost(ctx *gin.Context, c pb2.PostServiceClient) {
	body := &CreatPostRequestBody{}
	userId, exist := ctx.Get("userId")
	if !exist {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data := &pb2.PostData{
		Body:     body.Body,
		Head:     body.Head,
		Category: body.Category,
		Tags:     body.Tags,
	}

	res, err := c.CreateNewPost(context.Background(), &pb2.CreateNewPostRequest{
		UserId: userId.(int64),
		Data:   data,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}
