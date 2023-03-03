package post

import (
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/auth"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/post/routes"
)

func RegisterRouts(r *gin.Engine, postURL string, authService *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authService)

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
	routes.GetAllPosts(ctx, svc.Client)
}
func (svc *ServiceClient) getPostById(ctx *gin.Context) {
	routes.GetPostById(ctx, svc.Client)
}
func (svc *ServiceClient) createNewPost(ctx *gin.Context) {
	routes.CreateNewPost(ctx, svc.Client)
}
func (svc *ServiceClient) updatePostById(ctx *gin.Context) {

}
func (svc *ServiceClient) deletePostById(ctx *gin.Context) {

}
