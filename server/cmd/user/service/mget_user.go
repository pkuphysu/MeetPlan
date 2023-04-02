package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/user"
)

type MGetUserServiceI interface {
	MGetUser(ctx context.Context, req *user.MGetUserReq) ([]*user.User, error)
}

func NewMGetUserService(ctx context.Context) MGetUserServiceI {
	return &mGetUserService{}
}

type mGetUserService struct{}

func (s *mGetUserService) MGetUser(ctx context.Context, req *user.MGetUserReq) ([]*user.User, error) {
	dao := query.Q.WithContext(ctx).User
	if req.IsActive != nil {
		dao = dao.Where(query.Q.User.IsActive.Is(*req.IsActive))
	}
	if req.IsAdmin != nil {
		dao = dao.Where(query.Q.User.IsAdmin.Is(*req.IsAdmin))
	}
	if req.IsTeacher != nil {
		dao = dao.Where(query.Q.User.IsTeacher.Is(*req.IsTeacher))
	}
	if req.PageParam != nil {
		dao = dao.Limit(int(req.PageParam.PageSize)).Offset(int(req.PageParam.PageSize * (req.PageParam.PageNum - 1)))
	}

	users, err := dao.Find()
	if err != nil {
		return nil, err
	}
	return packUsers(users), nil
}
