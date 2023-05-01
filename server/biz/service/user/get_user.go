package user

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"

	"meetplan/biz/dal/pack"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type GetUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	UserDao        query.IUserDo
}

func NewGetUserService(ctx context.Context, RequestContext *app.RequestContext) *GetUserService {
	return &GetUserService{RequestContext: RequestContext, Context: ctx, UserDao: query.User.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *GetUserService) Run(req *model.GetUserRequest, resp *model.GetUserResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in GetUserService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.GetUserResponse)
	}

	user, e := h.UserDao.Where(query.User.ID.Eq(req.Id)).First()
	if e != nil && errors.Is(e, gorm.ErrRecordNotFound) {
		return errno.NewNotFoundErr("user not found")
	} else if e != nil {
		return errno.NewInternalErr("get user failed")
	}

	resp.Data = pack.UserDal2Vo(user)

	return
}
