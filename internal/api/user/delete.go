package user

import (
	"context"
	desc "github.com/MaksimovDenis/auth/pkg/userAPI_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	_, err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
