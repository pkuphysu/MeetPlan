package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/user"
	"github.com/pkuphysu/meetplan/pkg/errno"
)

type UpdateUserServiceI interface {
	UpdateUser(ctx context.Context, req *user.UpdateUserReq) error
}

func NewUpdateUserService(ctx context.Context) UpdateUserServiceI {
	return &updateUserService{}
}

type updateUserService struct{}

func (s *updateUserService) UpdateUser(ctx context.Context, req *user.UpdateUserReq) error {
	dao := query.Q.WithContext(ctx).User

	updateMap := map[string]interface{}{}
	if req.User.Id != nil {
		dao = dao.Where(query.Q.User.ID.Eq(*req.User.Id))
		if req.User.PkuId != nil {
			updateMap["pku_id"] = *req.User.PkuId
		}
	} else if req.User.PkuId != nil {
		dao = dao.Where(query.Q.User.PkuID.Eq(*req.User.PkuId))
	} else {
		return errno.ParamErr
	}

	if req.User.IsActive != nil {
		updateMap["is_active"] = *req.User.IsActive
	}
	if req.User.IsAdmin != nil {
		updateMap["is_admin"] = *req.User.IsAdmin
	}

	if len(updateMap) == 0 {
		return nil
	}
	res, err := dao.Updates(updateMap)
	if err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return errno.UserNotFoundErr
	}
	return nil
}
