package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type GetTermDateRangeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetTermDateRangeService(ctx context.Context, RequestContext *app.RequestContext) *GetTermDateRangeService {
	return &GetTermDateRangeService{RequestContext: RequestContext, Context: ctx}
}

func (h *GetTermDateRangeService) Run(req *model.GetTermDateRangeRequest, resp *model.GetTermDateRangeResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
