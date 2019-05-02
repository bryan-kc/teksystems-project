package v1

import (
	"github.com/boltdb/bolt"
	"context"

	"../../../pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

type blogServiceServer struct {
	db *bolt.DB
}

// NewToDoServiceServer creates  serviceToDo
func NewBlogServiceServer(db *bolt.DB) v1.BlogServiceServer {
	return &blogServiceServer{db: db}
}

func (blogServiceServer) ListPosts(ctx context.Context, req *v1.ListPostsRequest) (*v1.ListPostsResponse, error) {
	panic("implement me")
}

func (blogServiceServer) GetPost(ctx context.Context, req *v1.GetPostRequest) (*v1.GetPostResponse, error) {
	panic("implement me")
}

func (blogServiceServer) CreatePost(ctx context.Context, req *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	panic("implement me")
}

func (blogServiceServer) CreateComment(ctx context.Context, req *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	panic("implement me")
}

