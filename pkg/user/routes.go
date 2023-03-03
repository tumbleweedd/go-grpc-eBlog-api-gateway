package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/auth"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/user/routes"
)

func RegisterRoutes(r *gin.Engine, userURL string, authService *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authService)

	svc := &ServiceClient{
		Client: InitServiceClient(userURL),
	}

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.Use(a.AuthRequired)
			users.GET("/userList", svc.getUserList)
			users.GET("/me", svc.getLoggedUserProfile)
			users.GET("/:username/profile", svc.getUserProfile)
			users.GET("/:username/posts", svc.getUserPosts)
			users.POST("/", svc.addUser)
			users.PUT("/:username", svc.updateUserProfile)
			users.DELETE("/:username", svc.deleteUserProfile)
			//users.PUT("/:username/giveAdmin", svc.giveAdminRole)
			//users.PUT("/:username/takeAdmin", h.takeAdminRole)
		}

	}
}

func (svc *ServiceClient) getUserList(ctx *gin.Context) {
	routes.GetUserList(ctx, svc.Client)
}

func (svc *ServiceClient) getLoggedUserProfile(ctx *gin.Context) {
	routes.GetLoggedUserProfile(ctx, svc.Client)
}

func (svc *ServiceClient) getUserProfile(ctx *gin.Context) {
	routes.GetUserProfile(ctx, svc.Client)
}

func (svc *ServiceClient) addUser(ctx *gin.Context) {
	routes.AddUser(ctx, svc.Client)
}

func (svc *ServiceClient) updateUserProfile(ctx *gin.Context) {
	routes.UpdateUserProfile(ctx, svc.Client)
}

func (svc *ServiceClient) deleteUserProfile(ctx *gin.Context) {
	routes.DeleteUserProfile(ctx, svc.Client)
}

func (svc *ServiceClient) getUserPosts(ctx *gin.Context) {
	routes.GetUserPosts(ctx, svc.Client)
}
