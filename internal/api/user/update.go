package user

import (
	"context"
	"github.com/MaksimovDenis/auth/internal/converter"
	desc "github.com/MaksimovDenis/auth/pkg/userAPI_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	_, err := i.userService.Update(ctx, converter.ToUserUpdateFromDesc(req.GetUser()))
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
