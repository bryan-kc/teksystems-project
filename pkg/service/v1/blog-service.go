package v1

import (
	"context"
	"github.com/boltdb/bolt"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"math/rand"

	"../../../pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion  = "v1"
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	idLength    = 12
)

type blogServiceServer struct {
	db       *bolt.DB
	hashSeed int64
}

// NewToDoServiceServer creates  serviceToDo
func NewBlogServiceServer(db *bolt.DB, seed int64) v1.BlogServiceServer {
	return &blogServiceServer{db: db, hashSeed: seed}
}

func (s *blogServiceServer) ListPosts(ctx context.Context, req *v1.ListPostsRequest) (*v1.ListPostsResponse, error) {
	resp := &v1.ListPostsResponse{}

	posts := []*v1.Post{}

	err := s.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("posts")).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s", k, v)

			post := &v1.Post{}
			err := proto.Unmarshal(v, post)
			if err != nil {
				return err
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
	resp := &v1.GetPostResponse{}

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		respBytes := b.Get([]byte(req.Id))

		if respBytes == nil {
			// Some kind of nil error. 404
			return error()
		}

		err := proto.Unmarshal(respBytes, resp)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *blogServiceServer) CreatePost(ctx context.Context, req *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {

	postID := quickRandString(idLength)

	var comments []*v1.Comment
	post := &v1.Post{
		Author:   req.Author,
		Title:    req.Title,
		Text:     req.Text,
		Comments: comments,
	}

	err := s.db.View(func(tx *bolt.Tx) error {
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
		Id:   postID,
		Post: post,
	}, nil
}

func (s *blogServiceServer) CreateComment(ctx context.Context, req *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	post := &v1.Post{}
	comment := &v1.Comment{
		Author: req.Author,
		Text: req.Text,
	}

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		respBytes := b.Get([]byte(req.Id))

		if respBytes == nil {
			// Some kind of nil error. 404
			return error()
		}

		err := proto.Unmarshal(respBytes, post)
		if err != nil {
			return err
		}

		post.Comments = append(post.Comments, comment)

		postBytes, err := proto.Marshal(post)
		if err != nil {
			return err
		}

		err = b.Put([]byte(req.Id), postBytes)
		if err != nil {
			return err
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
