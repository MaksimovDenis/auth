package converter

import (
	"github.com/MaksimovDenis/auth/internal/repository/user/model"
	desc "github.com/MaksimovDenis/auth/pkg/userAPI_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToUserFromService(user *model.User) *desc.User {
	var updateAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updateAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      ToRoleFromService(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updateAt,
	}
}

func ToUserCreateFromService(userCreate *model.UserCreate) *desc.UserCreate {
	return &desc.UserCreate{
		Name:            userCreate.Name,
		Email:           userCreate.Email,
		Password:        userCreate.Password,
		PasswordConfirm: userCreate.PasswordConfirm,
		Role:            ToRoleFromService(userCreate.Role),
	}
}

func ToUserUpdateFromService(userUpdate *model.UserUpdate) *desc.UserUpdate {
	return &desc.UserUpdate{
		Id:    wrapperspb.Int64(userUpdate.ID),
		Name:  wrapperspb.String(userUpdate.Name),
		Email: wrapperspb.String(userUpdate.Email),
	}
}

func ToRoleFromService(role model.Role) desc.Role {
	switch role {
	case model.Role_USER:
		return desc.Role_USER
	case model.Role_ADMIN:
		return desc.Role_USER
	default:
		return desc.Role_USER
	}
}
