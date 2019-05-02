package cmd

import (
	"../protocol/grpc"
	"../protocol/rest"
	"context"
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	v1 "teksystems/pkg/service/v1"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string
}

func RunServer() error {
	ctx := context.Background()

	var cfg Config

	// Create new DB
	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	v1API := v1.NewBlogServiceServer(db)

	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
