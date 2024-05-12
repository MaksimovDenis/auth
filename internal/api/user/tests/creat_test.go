package tests

import (
	"context"
	"fmt"
	users "github.com/MaksimovDenis/auth/internal/api/user"
	"github.com/MaksimovDenis/auth/internal/model"
	"github.com/MaksimovDenis/auth/internal/service"
	serviceMocks "github.com/MaksimovDenis/auth/internal/service/mocks"
	desc "github.com/MaksimovDenis/auth/pkg/userAPI_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.BeerName()
		role     = 0

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			User: &desc.UserCreate{
				Name:     name,
				Email:    email,
				Password: password,
				Role:     desc.Role(role),
			},
		}

		user = &model.UserCreate{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     0,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, user).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, user).Return(0, serviceErr)
				return mock
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := users.NewImplementation(userServiceMock)

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)

		})
	}
}
