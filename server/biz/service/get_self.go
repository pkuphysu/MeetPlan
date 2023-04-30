package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type GetSelfService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetSelfService(ctx context.Context, RequestContext *app.RequestContext) *GetSelfService {
	return &GetSelfService{RequestContext: RequestContext, Context: ctx}
}

func (h *GetSelfService) Run(req *model.GetSelfRequest, resp *model.GetSelfResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
