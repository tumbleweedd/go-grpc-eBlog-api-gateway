package user

import (
	"fmt"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/internal/user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.UserServiceClient
}

func InitServiceClient(userSvcUrl string) pb.UserServiceClient {
	cc, err := grpc.Dial(userSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewUserServiceClient(cc)
}
