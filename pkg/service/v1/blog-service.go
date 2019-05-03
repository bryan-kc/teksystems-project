package v1

import (
	"context"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/bryan-kc/teksystems-project/pkg/api/v1"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	idLength    = 12
)

type blogServiceServer struct {
	db *bolt.DB
}

// NewToDoServiceServer creates  serviceToDo
func NewBlogServiceServer(db *bolt.DB) v1.BlogServiceServer {
	return &blogServiceServer{db: db}
}

func (s *blogServiceServer) ListPosts(ctx context.Context, req *v1.ListPostsRequest) (*v1.ListPostsResponse, error) {
	resp := &v1.ListPostsResponse{}

	posts := []*v1.Post{}

	err := s.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("posts")).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			post := &v1.Post{}
			err := proto.Unmarshal(v, post)
			if err != nil {
				return status.Error(codes.Unknown, fmt.Sprintf("Error Fetching Posts: %s", err.Error()))
			}

			posts = append(posts, post)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	resp.Posts = posts

	return resp, nil
}

func (s *blogServiceServer) GetPost(ctx context.Context, req *v1.GetPostRequest) (*v1.GetPostResponse, error) {
	resp := &v1.Post{}

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		respBytes := b.Get([]byte(req.Id))

		if respBytes == nil {
			return status.Error(codes.NotFound, fmt.Sprintf("Post not found: %s", req.Id))
		}

		err := proto.Unmarshal(respBytes, resp)
		if err != nil {
			return status.Error(codes.Unknown, fmt.Sprintf("Error Fetching Post: %s", err.Error()))
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &v1.GetPostResponse{Post: resp}, nil
}

func (s *blogServiceServer) CreatePost(ctx context.Context, req *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {

	postID := quickRandString(idLength)

	var comments []*v1.Comment
	post := &v1.Post{
		Id:       postID,
		Author:   req.Author,
		Title:    req.Title,
		Text:     req.Text,
		Comments: comments,
	}

	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))

		postBytes, err := proto.Marshal(post)
		if err != nil {
			return err
		}

		err = b.Put([]byte(postID), postBytes)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &v1.CreatePostResponse{
		Post: post,
	}, nil
}

func (s *blogServiceServer) CreateComment(ctx context.Context, req *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	post := &v1.Post{}
	comment := &v1.Comment{
		Author: req.Author,
		Text:   req.Text,
	}

	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		respBytes := b.Get([]byte(req.Id))

		if respBytes == nil {
			return status.Error(codes.NotFound, fmt.Sprintf("Post not found: %s", req.Id))
		}

		err := proto.Unmarshal(respBytes, post)
		if err != nil {
			return status.Error(codes.Unknown, fmt.Sprintf("Error Fetching Post: %s", err.Error()))
		}

		post.Comments = append(post.Comments, comment)

		postBytes, err := proto.Marshal(post)
		if err != nil {
			return status.Error(codes.Unknown, fmt.Sprintf("Error Writing Comment: %s", err.Error()))
		}

		err = b.Put([]byte(req.Id), postBytes)
		if err != nil {
			return status.Error(codes.Unknown, fmt.Sprintf("Error Writing Comment: %s", err.Error()))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &v1.CreateCommentResponse{
		Post: post,
	}, nil
}

func quickRandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}

	return string(b)
}
