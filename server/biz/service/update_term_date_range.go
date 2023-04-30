package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type UpdateTermDateRangeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateTermDateRangeService(ctx context.Context, RequestContext *app.RequestContext) *UpdateTermDateRangeService {
	return &UpdateTermDateRangeService{RequestContext: RequestContext, Context: ctx}
}

func (h *UpdateTermDateRangeService) Run(req *model.UpdateTermDateRangeRequest, resp *model.UpdateTermDateRangeResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
