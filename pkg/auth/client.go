package auth

import (
	"fmt"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/auth/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(authSvcUrl string) pb.AuthServiceClient {
	cc, err := grpc.Dial(authSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
