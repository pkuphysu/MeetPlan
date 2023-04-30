package service

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

func (h *GetUserService) Run(req *model.GetUserRequest, resp *model.GetUserResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
