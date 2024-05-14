package user

import (
	"context"
	"github.com/MaksimovDenis/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, create *model.UserCreate) (int64, error) {
	id, err := s.userRepository.Create(ctx, create)
	if err != nil {
		return 0, err
	}
	return id, nil
}
