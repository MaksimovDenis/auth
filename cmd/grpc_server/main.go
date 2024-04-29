package main

import (
	"context"
	"log"

	desc "github.com/MaksimovDenis/auth/pkg/userAPI_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserAPIV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())

}
