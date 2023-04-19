package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"meetplan/gorm_gen/query"
	"meetplan/kitex_gen/pkuphy/meetplan/user"
	"meetplan/pkg/errno"
)

type GetUserServiceI interface {
	GetUser(req *user.GetUserReq) (*user.User, error)
}

func NewGetUserService(ctx context.Context) GetUserServiceI {
	return &getUserService{
		ctx: ctx,
		dao: query.Q.WithContext(ctx).User,
	}
}

type getUserService struct {
	ctx context.Context
	dao query.IUserDo
}

func (s *getUserService) GetUser(req *user.GetUserReq) (*user.User, error) {
	dao := s.dao
	if req.Id != nil {
		dao = dao.Where(query.Q.User.ID.Eq(*req.Id))
	} else if req.PkuId != nil {
		dao = dao.Where(query.Q.User.PkuID.Eq(*req.PkuId))
	} else {
		return nil, errno.ParamErr.WithMessage("id_or_pku_id_is_required")
	}

	u, err := dao.First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errno.UserNotFoundErr
	} else if err != nil {
		return nil, err
	}
	return packUser(u), nil
}
