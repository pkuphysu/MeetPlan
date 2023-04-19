package service

import (
	"context"
	"meetplan/gorm_gen/query"
	"meetplan/kitex_gen/pkuphy/meetplan/user"
	"meetplan/pkg/errno"
)

type UpdateUserServiceI interface {
	UpdateUser(req *user.UpdateUserReq) error
}

func NewUpdateUserService(ctx context.Context) UpdateUserServiceI {
	return &updateUserService{
		ctx: ctx,
		dao: query.Q.WithContext(ctx).User,
	}
}

type updateUserService struct {
	ctx context.Context
	dao query.IUserDo
}

func (s *updateUserService) UpdateUser(req *user.UpdateUserReq) error {
	dao := s.dao

	updateMap := map[string]interface{}{}
	if req.User.Id != nil {
		dao = dao.Where(query.Q.User.ID.Eq(*req.User.Id))
		if req.User.PkuId != nil {
			updateMap["pku_id"] = *req.User.PkuId
		}
	} else if req.User.PkuId != nil {
		dao = dao.Where(query.Q.User.PkuID.Eq(*req.User.PkuId))
	} else {
		return errno.ParamErr.WithMessage("id_or_pku_id_is_required")
	}

	if req.User.Name != nil {
		updateMap["name"] = *req.User.Name
	}
	if req.User.Email != nil {
		updateMap["email"] = *req.User.Email
	}
	if req.User.IsActive != nil {
		updateMap["is_active"] = *req.User.IsActive
	}
	if req.User.IsAdmin != nil {
		updateMap["is_admin"] = *req.User.IsAdmin
	}
	if req.User.IsTeacher != nil {
		updateMap["is_teacher"] = *req.User.IsTeacher
	}
	if req.User.Gender != nil {
		updateMap["gender"] = *req.User.Gender
	}
	if req.User.Avatar != nil {
		updateMap["avatar"] = *req.User.Avatar
	}
	if req.User.Department != nil {
		updateMap["department"] = *req.User.Department
	}
	if req.User.Phone != nil {
		updateMap["phone"] = *req.User.Phone
	}
	if req.User.Major != nil {
		updateMap["major"] = *req.User.Major
	}
	if req.User.Grade != nil {
		updateMap["grade"] = *req.User.Grade
	}
	if req.User.Dorm != nil {
		updateMap["dorm"] = *req.User.Dorm
	}
	if req.User.Office != nil {
		updateMap["office"] = *req.User.Office
	}
	if req.User.Introduction != nil {
		updateMap["introduction"] = *req.User.Introduction
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
