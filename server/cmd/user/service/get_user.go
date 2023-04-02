package service

import (
	"context"
	"errors"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/user"
	"github.com/pkuphysu/meetplan/pkg/errno"
	"gorm.io/gorm"
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
