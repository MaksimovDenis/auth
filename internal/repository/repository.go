package repository

import (
	"context"
	"github.com/MaksimovDenis/auth/internal/model"
	"github.com/golang/protobuf/ptypes/empty"
)

type UserRepository interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, create *model.UserCreate) (int64, error)
	Update(ctx context.Context, update *model.UserUpdate) (*empty.Empty, error)
	Delete(ctx context.Context, id int64) (*empty.Empty, error)
}
