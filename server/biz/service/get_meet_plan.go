package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type GetMeetPlanService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetMeetPlanService(ctx context.Context, RequestContext *app.RequestContext) *GetMeetPlanService {
	return &GetMeetPlanService{RequestContext: RequestContext, Context: ctx}
}

func (h *GetMeetPlanService) Run(req *model.GetMeetPlanRequest, resp *model.GetMeetPlanResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
