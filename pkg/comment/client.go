package comment

import (
	"fmt"
	"github.com/tumbleweedd/grpc-eBlog/grpc-eBlog-api-gateway/pkg/comment/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.CommentServiceClient
}

func InitServiceClient(authSvcUrl string) pb.CommentServiceClient {
	cc, err := grpc.Dial(authSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCommentServiceClient(cc)
}
