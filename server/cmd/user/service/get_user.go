package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/user"
	"github.com/pkuphysu/meetplan/pkg/errno"
)

type GetUserServiceI interface {
	GetUser(ctx context.Context, req *user.GetUserReq) (*user.User, error)
}

func NewGetUserService(ctx context.Context) GetUserServiceI {
	return &getUserService{}
}

type getUserService struct{}

func (s *getUserService) GetUser(ctx context.Context, req *user.GetUserReq) (*user.User, error) {
	dao := query.Q.WithContext(ctx).User
	if req.Id != nil {
		dao = dao.Where(query.Q.User.ID.Eq(*req.Id))
	} else if req.PkuId != nil {
		dao = dao.Where(query.Q.User.PkuID.Eq(*req.PkuId))
	} else {
		return nil, errno.ParamErr
	}

	u, err := dao.First()
	if err != nil {
		return nil, err
	}
	return packUser(u), nil
}
