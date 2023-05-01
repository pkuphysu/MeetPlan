package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type GetUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserService(ctx context.Context, RequestContext *app.RequestContext) *GetUserService {
	return &GetUserService{RequestContext: RequestContext, Context: ctx}
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
	// todo edit your code
	return
}
