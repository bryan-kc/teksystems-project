pb:
	protoc -I. \
      -I$(GOPATH)/src \
      -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
      --proto_path=api/proto/v1 \
      --grpc-gateway_out=logtostderr=true:pkg/api/v1 \
      --go_out=plugins=grpc:pkg/api/v1 \
      --swagger_out=logtostderr=true:api/proto/v1 \
      blog-service.proto


deps:
	go get -u google.golang.org/grpc
	go get -u github.com/gogo/protobuf/protoc-gen-gogoslick
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/bryan-kc/teksystems-project/pkg/api/v1
	go get -u github.com/bryan-kc/teksystems-project/pkg/cmd
	go get -u github.com/bryan-kc/teksystems-project/pkg/protocol/grpc
    go get -u github.com/bryan-kc/teksystems-project/pkg/protocol/rest
    go get -u github.com/bryan-kc/teksystems-project/pkg/service/v1
    go get -u -t ./...

fmt:
	gofmt -s -w .

server:
	go build