FROM golang:1.12.4-alpine3.9

ENV PROTOBUF_VERSION=3.6.1 \
    PROTOC_GEN_GO_VERSION=1.2.0 \
    PROTOC_GEN_GRPC_GATEWAY_VERSION=1.5.1

RUN apk add --no-cache \
    build-base \
    curl \
    automake \
    autoconf \
    libtool \
    git \
    zlib-dev

RUN mkdir -p /protobuf && \
    curl -L https://github.com/protocolbuffers/protobuf/archive/v${PROTOBUF_VERSION}.tar.gz | tar xvz --strip-components=1 -C /protobuf
RUN cd /protobuf && \
    ./autogen.sh && \
    ./configure --prefix=/usr && \
    make -j2 && make install

RUN apk add --no-cache go
RUN go get -u -v -ldflags '-w -s' \
    github.com/golang/protobuf/protoc-gen-go \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

COPY . /go/src/github.com/bryan-kc/teksystems-project

RUN go install github.com/bryan-kc/teksystems-project/pkg/cmd/server

#RUN go get -d -v ./...
#RUN go install -v ./...
ENTRYPOINT ["/go/bin/server"]
