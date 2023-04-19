package service

import (
	"context"
	"meetplan/gorm_gen/query"
	"meetplan/kitex_gen/pkuphy/meetplan/base"
	"meetplan/kitex_gen/pkuphy/meetplan/user"
)

type QueryUserServiceI interface {
	QueryUser(req *user.QueryUserReq) ([]*user.User, *base.PageParam, error)
}

func NewQueryUserService(ctx context.Context) QueryUserServiceI {
	return &queryUserService{
		ctx: ctx,
		dao: query.Q.WithContext(ctx).User,
	}
}

type queryUserService struct {
	ctx context.Context
	dao query.IUserDo
}

func (s *queryUserService) QueryUser(req *user.QueryUserReq) ([]*user.User, *base.PageParam, error) {
	dao := s.dao
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
