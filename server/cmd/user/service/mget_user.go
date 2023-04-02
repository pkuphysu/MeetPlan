package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/user"
	"github.com/pkuphysu/meetplan/pkg/errno"
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

	if len(req.Ids) == 0 && len(req.PkuIds) == 0 {
		return nil, errno.ParamErr.WithMessage("ids_or_pku_ids_is_required")
	} else if len(req.Ids) != 0 && len(req.Ids) != 0 {
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
