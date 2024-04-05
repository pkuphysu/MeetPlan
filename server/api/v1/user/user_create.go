package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type CreateUserRequest struct {
	Users []*model.User `json:"users"`
}

func CreateUsers(ctx context.Context, c *app.RequestContext, req *CreateUserRequest) ([]*model.User, *types.PageInfo, error) {
	err := query.UserColl.InsertMany(ctx, req.Users)
	if err != nil {
		return nil, nil, err
	}
	return req.Users, nil, nil
}
