package tests

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/MaksimovDenis/auth/internal/model"
	"github.com/MaksimovDenis/auth/internal/repository"
	serviceMocks "github.com/MaksimovDenis/auth/internal/service/mocks"
	"github.com/MaksimovDenis/auth/internal/service/user"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		role      = model.Role_USER
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		repoErr = fmt.Errorf("repo error")

		res = &model.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Role:      role,
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *model.User
		wantErr         error
		userServiceMock userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want:    res,
			wantErr: nil,
			userServiceMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(res, nil)
				return mock
			},
		},
		{
			name: "error case service",
			args: args{
				ctx: ctx,
				req: id,
			},
			want:    nil,
			wantErr: repoErr,
			userServiceMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, repoErr)
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			service := user.NewUserService(userServiceMock)

			newID, err := service.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, newID)
			require.Equal(t, tt.wantErr, err)
		})
	}

}
