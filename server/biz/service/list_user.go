package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type ListUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListUserService(ctx context.Context, RequestContext *app.RequestContext) *ListUserService {
	return &ListUserService{RequestContext: RequestContext, Context: ctx}
}

// Run req should not be nil and resp should not be nil
func (h *ListUserService) Run(req *model.ListUserRequest, resp *model.ListUserResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.ListUserResponse)
	}
	// todo edit your code
	return
}
