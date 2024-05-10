package user

import (
	"context"
	"github.com/MaksimovDenis/auth/internal/converter"
	desc "github.com/MaksimovDenis/auth/pkg/userAPI_v1"
	"log"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserCreateFromDesc(req.GetUser()))
	if err != nil {
		return nil, err
	}

	log.Printf("insertded note with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
