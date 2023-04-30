package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type ListMeetPlanService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListMeetPlanService(ctx context.Context, RequestContext *app.RequestContext) *ListMeetPlanService {
	return &ListMeetPlanService{RequestContext: RequestContext, Context: ctx}
}

func (h *ListMeetPlanService) Run(req *model.ListMeetPlanRequest, resp *model.ListMeetPlanResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
