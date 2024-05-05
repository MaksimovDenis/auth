package user

import desc "github.com/MaksimovDenis/auth/pkg/userAPI_v1"

func (i *Implemenation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserCreateFromSe)
}
