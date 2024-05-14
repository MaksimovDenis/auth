package tests

import (
	"context"
	"fmt"
	"github.com/MaksimovDenis/auth/internal/repository"
	repoMocks "github.com/MaksimovDenis/auth/internal/repository/mocks"
	"github.com/MaksimovDenis/auth/internal/service/user"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = gofakeit.Int64()

		repoErr = fmt.Errorf("repo error")

		req = id

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
				mock.DeleteMock.Expect(ctx, id).Return(res, nil)
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
				mock.DeleteMock.Expect(ctx, id).Return(nil, repoErr)
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

			newID, err := service.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.want, newID)
			require.Equal(t, tt.err, err)
		})
	}
}
