package user

import (
	"context"
	"github.com/MaksimovDenis/auth/internal/model"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *serv) Update(ctx context.Context, update *model.UserUpdate) (*empty.Empty, error) {
	_, err := s.userRepository.Update(ctx, update)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
