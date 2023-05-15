package user

import (
	"github.com/gin-gonic/gin"
	auth2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/auth"
	routes2 "github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/user/routes"
)

func RegisterRoutes(r *gin.Engine, userURL string, authService *auth2.ServiceClient) {
	a := auth2.InitAuthMiddleware(authService)

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
	routes2.GetUserList(ctx, svc.Client)
}

func (svc *ServiceClient) getLoggedUserProfile(ctx *gin.Context) {
	routes2.GetLoggedUserProfile(ctx, svc.Client)
}

func (svc *ServiceClient) getUserProfile(ctx *gin.Context) {
	routes2.GetUserProfile(ctx, svc.Client)
}

func (svc *ServiceClient) addUser(ctx *gin.Context) {
	routes2.AddUser(ctx, svc.Client)
}

func (svc *ServiceClient) updateUserProfile(ctx *gin.Context) {
	routes2.UpdateUserProfile(ctx, svc.Client)
}

func (svc *ServiceClient) deleteUserProfile(ctx *gin.Context) {
	routes2.DeleteUserProfile(ctx, svc.Client)
}

func (svc *ServiceClient) getUserPosts(ctx *gin.Context) {
	routes2.GetUserPosts(ctx, svc.Client)
}
