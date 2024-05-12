package converter

import (
	"github.com/MaksimovDenis/auth/internal/model"
	modelRepo "github.com/MaksimovDenis/auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      ToRoleFromRepo(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserCreateFromRepo(userCreate *modelRepo.UserCreate) *model.UserCreate {
	return &model.UserCreate{
		Name:            userCreate.Name,
		Email:           userCreate.Email,
		Password:        userCreate.Password,
		PasswordConfirm: userCreate.PasswordConfirm,
		Role:            ToRoleFromRepo(userCreate.Role),
	}
}

func ToUserUpdateFromRepo(userUpdate *modelRepo.UserUpdate) *model.UserUpdate {
	return &model.UserUpdate{
		ID:    userUpdate.ID,
		Name:  userUpdate.Name,
		Email: userUpdate.Email,
	}
}

func ToRoleFromRepo(role modelRepo.Role) model.Role {
	switch role {
	case modelRepo.Role_USER:
		return model.Role_USER
	case modelRepo.Role_ADMIN:
		return model.Role_ADMIN
	default:
		return model.Role_USER
	}
}
