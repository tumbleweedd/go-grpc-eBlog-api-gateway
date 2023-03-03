package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/post/pb"
	"net/http"
)

type CreatPostRequestBody struct {
	Body     string   `json:"body"`
	Head     string   `json:"title"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

func CreateNewPost(ctx *gin.Context, c pb.PostServiceClient) {
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

	data := &pb.PostData{
		Body:     body.Body,
		Head:     body.Head,
		Category: body.Category,
		Tags:     body.Tags,
	}

	res, err := c.CreateNewPost(context.Background(), &pb.CreateNewPostRequest{
		UserId: userId.(int64),
		Data:   data,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)

}
