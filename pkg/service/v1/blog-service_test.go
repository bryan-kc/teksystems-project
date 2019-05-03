package v1

import (
	"context"
	"github.com/boltdb/bolt"
	"github.com/bryan-kc/teksystems-project/pkg/api/v1"
	"log"
	"reflect"
	"testing"
)

func TestServerSetup(t *testing.T) {
	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("posts"))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		t.Fatalf("An error occurred while setting up the database: %s", err)
	}

	defer db.Close()

	server := NewBlogServiceServer(db)

	if server == nil {
		t.Fatalf("Failed to setup Blog Service Server")
	}

	// This test needs to be updated to use mock storage
	// testOrgBasics(t, ms)
	testCreatePost(t, server)
	testCreatePostandComment(t, server)
	testListPosts(t, server)
}

func testCreatePost(t *testing.T, server v1.BlogServiceServer) {
	ctx := context.Background()

	postReq := &v1.CreatePostRequest{
		Author: "Testy Testerson",
		Title:  "Testing Simulator 2019",
		Text:   "This game rocks!",
	}

	postResp, err := server.CreatePost(ctx, postReq)

	if err != nil {
		t.Errorf("Error in Creating Post: %s", err)
	}

	getPostReq := &v1.GetPostRequest{Id: postResp.Post.Id}
	getPostResp, err := server.GetPost(ctx, getPostReq)
	if err != nil {
		t.Errorf("Error in Fetching Post: %s", err)
	}

	if reflect.DeepEqual(postResp, getPostResp) {
		t.Errorf("Created post and returned post are not equal: [%+v], [%+v]", postResp, getPostResp)
	}
}

func testCreatePostandComment(t *testing.T, server v1.BlogServiceServer) {
	ctx := context.Background()

	postReq := &v1.CreatePostRequest{
		Author: "Testy Testerson",
		Title:  "Testing Simulator 2019",
		Text:   "This game rocks!",
	}

	postResp, err := server.CreatePost(ctx, postReq)

	if err != nil {
		t.Errorf("Error Creating Post: %s", err)
	}

	commentReq := &v1.CreateCommentRequest{
		Id:     postResp.Post.Id,
		Author: "GameNerd",
		Text:   "I agree this game is awesome",
	}

	commentResp, err := server.CreateComment(ctx, commentReq)
	if err != nil {
		t.Errorf("Error Creating Comment: %s", err)
	}

	getPostReq := &v1.GetPostRequest{Id: postResp.Post.Id}
	getPostResp, err := server.GetPost(ctx, getPostReq)
	if err != nil {
		t.Errorf("Error in Fetching Post: %s", err)
	}

	if !reflect.DeepEqual(commentResp.Post, getPostResp.Post) {
		t.Errorf("Created post and returned post are not equal: [%+v], [%+v]", postResp, getPostResp)
	}
}

func testListPosts(t *testing.T, server v1.BlogServiceServer) {
	ctx := context.Background()

	postReq1 := &v1.CreatePostRequest{
		Author: "Testy Tester",
		Title:  "Testing Simulator 2010",
		Text:   "This game rocks!",
	}

	postReq2 := &v1.CreatePostRequest{
		Author: "Testy Tester",
		Title:  "Testing Simulator 2011",
		Text:   "This game rocks!",
	}

	postReq3 := &v1.CreatePostRequest{
		Author: "Testy Tester",
		Title:  "Testing Simulator 2012",
		Text:   "This game rocks!",
	}

	_, err := server.CreatePost(ctx, postReq1)
	if err != nil {
		t.Errorf("Error Creating Post 1: %s", err)
	}

	_, err = server.CreatePost(ctx, postReq2)
	if err != nil {
		t.Errorf("Error Creating Post 2: %s", err)
	}

	_, err = server.CreatePost(ctx, postReq3)
	if err != nil {
		t.Errorf("Error Creating Post 3: %s", err)
	}

	listReq := &v1.ListPostsRequest{}
	listPostResp, err := server.ListPosts(ctx, listReq)

	// Note, posts already exist in database from previous tests.
	if len(listPostResp.Posts) != 5 {
		t.Errorf("Post List does not contain enough elements (%v): %+v", len(listPostResp.Posts), listPostResp)
	}
}
