package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/auth"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/post"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/user"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	r := gin.Default()

	authService := *auth.RegisterRoutes(r, viper.GetString("AUTH_SVC_URL"))
	user.RegisterRoutes(r, viper.GetString("USER_SVC_URL"), &authService)
	post.RegisterRouts(r, viper.GetString("POST_SVC_URL"), &authService)

	r.Run(viper.GetString("PORT"))
}

func initConfig() error {
	viper.AddConfigPath("pkg/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
