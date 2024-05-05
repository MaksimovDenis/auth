package user

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *serv) Delete(ctx context.Context, id int64) (*empty.Empty, error) {
	_, err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
