package user

import (
	"github.com/MaksimovDenis/auth/internal/repository"
	"github.com/MaksimovDenis/auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

func NewService(
	userRepository repository.UserRepository,
) service.UserService {
	return &serv{
		userRepository: userRepository,
	}
}

func NewUserService(deps ...interface{}) service.UserService {
	srv := serv{}

	for _, v := range deps {
		switch s := v.(type) {
		case repository.UserRepository:
			srv.userRepository = s
		}
	}
	return &srv
}
