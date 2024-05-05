package app

import (
	"github.com/MaksimovDenis/auth/internal/config"
	"github.com/MaksimovDenis/auth/internal/repository"
	"github.com/MaksimovDenis/auth/internal/service"
)

type ServiceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	userRepository repository.UserRepository

	userService service.UserService

	userImpl *user.Implementation
}
