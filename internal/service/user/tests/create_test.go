package tests

import (
	"context"
	"fmt"
	"github.com/MaksimovDenis/auth/internal/model"
	"github.com/MaksimovDenis/auth/internal/repository"
	repoMocks "github.com/MaksimovDenis/auth/internal/repository/mocks"
	"github.com/MaksimovDenis/auth/internal/service/user"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req *model.UserCreate
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.BeerName()
		role     = model.Role_USER

		repoErr = fmt.Errorf("repo error")

		req = &model.UserCreate{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		}
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(id, nil)
				return mock
			},
		},
		{
			name: "service case error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(0, repoErr)
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := tt.userRepositoryMock(mc)
			service := user.NewUserService(userRepoMock)

			newID, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, newID)
			require.Equal(t, tt.err, err)
		})
	}
}
