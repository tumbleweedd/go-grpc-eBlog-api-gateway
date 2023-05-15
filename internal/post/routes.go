package post

import (
	"github.com/gin-gonic/gin"
	auth2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/auth"
	routes2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/post/routes"
)

func RegisterRouts(r *gin.Engine, postURL string, authService *auth2.ServiceClient) {
	a := auth2.InitAuthMiddleware(authService)

	svc := &ServiceClient{
		Client: InitServiceClient(postURL),
	}

	api := r.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.Use(a.AuthRequired)
			posts.GET("/", svc.getAllPosts)
			posts.GET("/:id", svc.getPostById)
			posts.POST("/", svc.createNewPost)
			posts.PUT("/:id", svc.updatePostById)
			posts.DELETE("/:id", svc.deletePostById)
		}
	}
}

func (svc *ServiceClient) getAllPosts(ctx *gin.Context) {
	routes2.GetAllPosts(ctx, svc.Client)
}
func (svc *ServiceClient) getPostById(ctx *gin.Context) {
	routes2.GetPostById(ctx, svc.Client)
}
func (svc *ServiceClient) createNewPost(ctx *gin.Context) {
	routes2.CreateNewPost(ctx, svc.Client)
}
func (svc *ServiceClient) updatePostById(ctx *gin.Context) {

}
func (svc *ServiceClient) deletePostById(ctx *gin.Context) {
	routes2.DeletePostById(ctx, svc.Client)
}
