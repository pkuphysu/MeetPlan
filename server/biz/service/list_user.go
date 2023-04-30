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

func (h *ListUserService) Run(req *model.ListUserRequest, resp *model.ListUserResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
