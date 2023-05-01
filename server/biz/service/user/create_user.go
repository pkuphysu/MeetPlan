package user

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/samber/lo"

	"meetplan/biz/dal/pack"
	"meetplan/biz/gorm_gen"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type CreateUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	UserDAO        query.IUserDo
}

func NewCreateUserService(ctx context.Context, RequestContext *app.RequestContext) *CreateUserService {
	return &CreateUserService{RequestContext: RequestContext, Context: ctx, UserDAO: query.User.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *CreateUserService) Run(req *model.CreateUserRequest, resp *model.CreateUserResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateUserService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.CreateUserResponse)
	}

	user := &gorm_gen.User{
		ID:          0,
		PkuID:       req.PkuId,
		Name:        req.Name,
		Email:       req.Email,
		EmailChange: nil,
		IsActive:    true,
		IsTeacher: lo.IfF(req.IsTeacher != nil, func() bool {
			return *req.IsTeacher
		}).Else(false),
		IsAdmin: lo.IfF(req.IsAdmin != nil, func() bool {
			return *req.IsAdmin
		}).Else(false),
		Gender: lo.IfF(req.Gender != nil, func() *bool {
			return lo.ToPtr(*req.Gender == model.Gender_GENDER_FEMALE)
		}).Else(nil),
		Avatar:       req.Avatar,
		Department:   req.Department,
		Phone:        req.Phone,
		Major:        req.Major,
		Grade:        req.Grade,
		Dorm:         req.Dorm,
		Office:       req.Office,
		Introduction: req.Introduction,
	}

	e := h.UserDAO.Create(user)
	if e != nil {
		if strings.Contains(e.Error(), "Duplicate entry") {
			return errno.NewValidationErr("user already exists")
		}
		return errno.ToInternalErr(e)
	}

	resp.Data = pack.UserDal2Vo(user)

	return
}
