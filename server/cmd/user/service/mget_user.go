package service

import (
	"context"
	"meetplan/gorm_gen/query"
	"meetplan/kitex_gen/pkuphy/meetplan/user"
	"meetplan/pkg/errno"
)

type MGetUserServiceI interface {
	MGetUser(req *user.MGetUserReq) ([]*user.User, error)
}

func NewMGetUserService(ctx context.Context) MGetUserServiceI {
	return &mGetUserService{
		ctx: ctx,
		dao: query.Q.WithContext(ctx).User,
	}
}

type mGetUserService struct {
	ctx context.Context
	dao query.IUserDo
}

func (s *mGetUserService) MGetUser(req *user.MGetUserReq) ([]*user.User, error) {
	dao := s.dao

	if len(req.Ids) == 0 && len(req.PkuIds) == 0 {
		return nil, errno.ParamErr.WithMessage("ids_or_pku_ids_is_required")
	} else if len(req.Ids) != 0 && len(req.PkuIds) != 0 {
		dao = dao.Where(query.Q.User.ID.In(req.Ids...)).Or(query.Q.User.PkuID.In(req.PkuIds...))
	} else if len(req.Ids) != 0 {
		dao = dao.Where(query.Q.User.ID.In(req.Ids...))
	} else {
		dao = dao.Where(query.Q.User.PkuID.In(req.PkuIds...))
	}

	u, err := dao.Find()
	if err != nil {
		return nil, err
	}
	return packUsers(u), nil
}
