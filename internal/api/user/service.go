package user

import (
	"github.com/MaksimovDenis/auth/internal/service"
	desc "github.com/MaksimovDenis/auth/pkg/userAPI_v1"
)

type Implemenation struct {
	desc.UnimplementedUserAPIV1Server
	userService service.UserService
}

func NewImplemention(userService service.UserService) *Implemenation {
	return &Implemenation{
		userService: userService,
	}
}
