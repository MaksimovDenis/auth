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
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req *model.UserUpdate
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Email()
		email = gofakeit.Email()

		repoErr = fmt.Errorf("repo error")

		req = &model.UserUpdate{
			ID:    id,
			Name:  name,
			Email: email,
		}

		res = &empty.Empty{}
	)

	tests := []struct {
		name           string
		args           args
		want           *empty.Empty
		err            error
		userRepository userRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userRepository: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, req).Return(res, nil)
				return mock
			},
		},
		{
			name: "service case error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  repoErr,
			userRepository: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, req).Return(nil, repoErr)
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := tt.userRepository(mc)
			service := user.NewUserService(userRepoMock)

			newID, err := service.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, newID)
			require.Equal(t, tt.err, err)
		})
	}
}
