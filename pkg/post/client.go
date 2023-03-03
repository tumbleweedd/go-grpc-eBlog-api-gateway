package post

import (
	"fmt"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/post/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.PostServiceClient
}

func InitServiceClient(postSvcUrl string) pb.PostServiceClient {
	cc, err := grpc.Dial(postSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewPostServiceClient(cc)
}
