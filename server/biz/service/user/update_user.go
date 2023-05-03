package user

import (
	"context"
	"errors"

	"meetplan/biz/dal/pack"

	"gorm.io/gen/field"
	"gorm.io/gorm"

	"meetplan/biz/gorm_gen/query"

	"github.com/cloudwego/hertz/pkg/app"

	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type UpdateUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	UserDAO        query.IUserDo
}

func NewUpdateUserService(ctx context.Context, RequestContext *app.RequestContext) *UpdateUserService {
	return &UpdateUserService{RequestContext: RequestContext, Context: ctx, UserDAO: query.User.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *UpdateUserService) Run(req *model.UpdateUserRequest, resp *model.UpdateUserResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in UpdateUserService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.UpdateUserResponse)
	}

	oriUser, e := h.UserDAO.Where(query.User.ID.Eq(req.Id)).First()
	if e != nil && errors.Is(e, gorm.ErrRecordNotFound) {
		return errno.NewNotFoundErr("user not found")
	} else if e != nil {
		return errno.ToInternalErr(e)
	}

	var updates []field.AssignExpr
	if req.PkuId != nil && *req.PkuId != oriUser.PkuID {
		updates = append(updates, query.User.PkuID.Value(*req.PkuId))
	}
	if req.Name != nil && *req.Name != oriUser.Name {
		updates = append(updates, query.User.Name.Value(*req.Name))
	}
	if req.Email != nil && *req.Email != oriUser.Email {
		updates = append(updates, query.User.EmailChange.Value(*req.Email))
	}
	if req.IsTeacher != nil && *req.IsTeacher != oriUser.IsTeacher {
		updates = append(updates, query.User.IsTeacher.Value(*req.IsTeacher))
	}
	if req.IsAdmin != nil && *req.IsAdmin != oriUser.IsAdmin {
		updates = append(updates, query.User.IsAdmin.Value(*req.IsAdmin))
	}
	if req.IsActive != nil && *req.IsActive != oriUser.IsActive {
		updates = append(updates, query.User.IsActive.Value(*req.IsActive))
	}
	if req.Gender != nil {
		gender := *req.Gender == model.Gender_GENDER_FEMALE
		if oriUser.Gender == nil || *oriUser.Gender != gender {
			updates = append(updates, query.User.Gender.Value(gender))
		}
	}
	if req.Avatar != nil {
		if oriUser.Avatar == nil || *oriUser.Avatar != *req.Avatar {
			updates = append(updates, query.User.Avatar.Value(*req.Avatar))
		}
	}
	if req.Phone != nil {
		if oriUser.Phone == nil || *oriUser.Phone != *req.Phone {
			updates = append(updates, query.User.Phone.Value(*req.Phone))
		}
	}
	if req.Department != nil {
		if oriUser.Department == nil || *oriUser.Department != *req.Department {
			updates = append(updates, query.User.Department.Value(*req.Department))
		}
	}
	if req.Major != nil {
		if oriUser.Major == nil || *oriUser.Major != *req.Major {
			updates = append(updates, query.User.Major.Value(*req.Major))
		}
	}
	if req.Grade != nil {
		if oriUser.Grade == nil || *oriUser.Grade != *req.Grade {
			updates = append(updates, query.User.Grade.Value(*req.Grade))
		}
	}
	if req.Dorm != nil {
		if oriUser.Dorm == nil || *oriUser.Dorm != *req.Dorm {
			updates = append(updates, query.User.Dorm.Value(*req.Dorm))
		}
	}
	if req.Office != nil {
		if oriUser.Office == nil || *oriUser.Office != *req.Office {
			updates = append(updates, query.User.Office.Value(*req.Office))
		}
	}
	if req.Introduction != nil {
		if oriUser.Introduction == nil || *oriUser.Introduction != *req.Introduction {
			updates = append(updates, query.User.Introduction.Value(*req.Introduction))
		}
	}
	if len(updates) == 0 {
		resp.Data = pack.UserDal2Vo(oriUser)
		return
	}

	_, e = h.UserDAO.Where(query.User.ID.Eq(req.Id)).UpdateColumnSimple(updates...)
	if e != nil {
		return errno.ToInternalErr(e)
	}
	newUser, e := h.UserDAO.Where(query.User.ID.Eq(req.Id)).First()
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = pack.UserDal2Vo(newUser)
	return
}
