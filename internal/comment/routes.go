package comment

import (
	"github.com/gin-gonic/gin"
	auth2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/auth"
	routes2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/comment/routes"
)

func RegisterRoutes(r *gin.Engine, commentURL string, authService *auth2.ServiceClient) {
	a := auth2.InitAuthMiddleware(authService)

	svc := &ServiceClient{
		Client: InitServiceClient(commentURL),
	}

	api := r.Group("/api")
	{
		posts := api.Group("/posts/:id")
		{
			comments := posts.Group("/comments")
			{
				comments.Use(a.AuthRequired)
				comments.POST("/", svc.addComment)
				comments.GET("/", svc.getCommentsByPostId)
				comments.GET("/:commentId", svc.getCommentByCommentId)
				comments.DELETE("/:commentId", svc.deleteCommentOnPostByCommentId)
			}
		}
	}
}

func (svc *ServiceClient) addComment(ctx *gin.Context) {
	routes2.AddComment(ctx, svc.Client)
}

func (svc *ServiceClient) getCommentsByPostId(ctx *gin.Context) {
	routes2.GetCommentsByPostId(ctx, svc.Client)
}

func (svc *ServiceClient) getCommentByCommentId(ctx *gin.Context) {
	routes2.GetCommentByCommentId(ctx, svc.Client)
}

func (svc *ServiceClient) deleteCommentOnPostByCommentId(ctx *gin.Context) {
	routes2.DeleteCommentOnPostByCommentId(ctx, svc.Client)
}
