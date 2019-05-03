package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type comment struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

type post struct {
	ID string `json:"id"`
	Author string `json:"author"`
	Title string `json:"title"`
	Text string `json:"text"`
	Comments []string `json:"comments"`
}

func main() {
	// get configuration
	address := flag.String("server", "http://localhost:8080", "HTTP gateway url, e.g. http://localhost:8080")
	flag.Parse()

	var body string

	// Call Create
	resp, err := http.Post(*address+"/v1/post", "application/json", strings.NewReader(fmt.Sprintf(`
		{"author": "Testerson",
		"title": "Testing Simulator 2019",
		"text": "This game sucks!"
		}
	`)))
	if err != nil {
		log.Fatalf("failed to call Create method: %v", err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Create response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Create response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	// parse ID of created Post
	var postResp struct {
		Post     post `json:"post"`
	}
	err = json.Unmarshal(bodyBytes, &postResp)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON response of Create method: %v", err)
		fmt.Println("error:", err)
	}

	// Call Get Post
	resp, err = http.Get(fmt.Sprintf("%s/%s/", *address, fmt.Sprintf("v1/post/%s", postResp.Post.ID)))
	if err != nil {
		log.Fatalf("failed to call Read method: %v", err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Read response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Read response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	// Call Create Comment
	resp, err = http.Post(fmt.Sprintf("%s/%s/", *address, fmt.Sprintf("v1/post/%s", postResp.Post.ID)), "application/json", strings.NewReader(fmt.Sprintf(`
		{"id", "%s",
		"author": "Bryan",
		"text": "I agree"
		}
	`, postResp.Post.ID)))

	if err != nil {
		log.Fatalf("failed to call Create comment method: %v", err)
	}

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Create comment response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Create comment response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	var postWithComment struct {
		ID       string    `json:"id"`
		Author   string    `json:"author"`
		Text     string    `json:"text"`
		Comments []comment `json:"comments"`
	}
	err = json.Unmarshal(bodyBytes, &postWithComment)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON response of Create method: %v", err)
		fmt.Println("error:", err)
	}
}
