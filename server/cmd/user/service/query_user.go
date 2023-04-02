package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/base"
	"github.com/pkuphysu/meetplan/kitex_gen/user"
)

type QueryUserServiceI interface {
	QueryUser(ctx context.Context, req *user.QueryUserReq) ([]*user.User, *base.PageParam, error)
}

func NewQueryUserService(ctx context.Context) QueryUserServiceI {
	return &queryUserService{}
}

type queryUserService struct{}

func (s *queryUserService) QueryUser(ctx context.Context, req *user.QueryUserReq) ([]*user.User, *base.PageParam, error) {
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

	var pageParam *base.PageParam
	if req.PageParam != nil {
		pageParam = req.PageParam
	} else {
		pageParam = base.NewPageParam()
		pageParam.SetPageNum(1)
		pageParam.SetPageSize(10)
	}
	dao = dao.Limit(int(pageParam.PageSize)).Offset(int(pageParam.PageSize * (pageParam.PageNum - 1)))

	users, err := dao.Find()
	if err != nil {
		return nil, nil, err
	}
	return packUsers(users), pageParam, nil
}
