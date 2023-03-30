package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/auth"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/comment/routes"
)

func RegisterRoutes(r *gin.Engine, commentURL string, authService *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authService)

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
	routes.AddComment(ctx, svc.Client)
}

func (svc *ServiceClient) getCommentsByPostId(ctx *gin.Context) {
	routes.GetCommentsByPostId(ctx, svc.Client)
}

func (svc *ServiceClient) getCommentByCommentId(ctx *gin.Context) {
	routes.GetCommentByCommentId(ctx, svc.Client)
}

func (svc *ServiceClient) deleteCommentOnPostByCommentId(ctx *gin.Context) {
	routes.DeleteCommentOnPostByCommentId(ctx, svc.Client)
}
