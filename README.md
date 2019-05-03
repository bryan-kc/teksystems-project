# TEKsystems DevOps Project

## Requirements

- Built on Go `1.12.4`
- Utilizes Protocol Buffers for gRPC/HTTP endpoints https://github.com/golang/protobuf
- Utilized extended grpc ecosystem https://github.com/grpc-ecosystem/grpc-gateway
- Utilized BoltDB for in-memory object storage https://github.com/boltdb/bolt

## User Story
As an avid video game reviewer
I want a way to create blog posts for my video game reviews
So that I can share my reviews in a way that my readers can respond to.

### Breakdown
- Single user (video game reviewer)
- Way to create blog posts for video game reviews
    - What is a blog post?
    - What makes up a 'video game review?'
        - Title, text discussion and author
- So that I can share my reviews in a way that my readers can respond to
    - _Share my reviews_
        - Sharing is inherent with post creation
    - "In a way my readers can respond to"
        - Implies a comment section

### Acceptance Criteria
- A blog post will show a title, article text (plain text) and an author name
- Comments are made on blog posts and show comment text (plain text) and an author name

## Development
The project's structure was designed such that the API routs are easily extensible.

To develop on the project:
1. Update the `*.proto` file in the `api/proto/v1` folder to add new messages and endpoints
2. At the root of the project run `make pb` to generate the protobuf files and the associated application gateway
3. Update `pkg/service/blog-service.go` to utilize the new endpoints and message types
4. Update `pkg/service/blog-service_test.go` to test the new endpoints.

## Usage


## Testing
- The test of the service can be run by running `go test` inside of `pkg/api/v1/service`
    - Note. This method of testing only utilizes gRPC application endpoints. 

## API Routes
API Route documentation in the form of a swagger can be found (here) [here](api/proto/v1/blog-service.swagger.json)

## Discussion
- We assume that the API can only be accessed by authenticated users. 
At the moment there is no authentication on any API routes.
- By default, the maximum message size to send/receive over gRPC is 4MB. 
We assume that at the moment the with the current storage that posts and their comments will be less than 4MB.
    - This can be alleviated in the future by increasing the gRPC message size
    - This can also be fixed by decoupling comments and posts, such that fetching a post is a particular call, and fetching a posts' comments is a separate call

## Future Goals
- Decouple comments from the posts themselves, and have each comment link to a postID
    - Comments can also store timestamps.
    - Comments could also have replies.
- Posts could have a score field. Either a rating system (stars) or a score out of 100.

