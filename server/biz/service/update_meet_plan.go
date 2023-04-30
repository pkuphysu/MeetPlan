package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type UpdateMeetPlanService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateMeetPlanService(ctx context.Context, RequestContext *app.RequestContext) *UpdateMeetPlanService {
	return &UpdateMeetPlanService{RequestContext: RequestContext, Context: ctx}
}

func (h *UpdateMeetPlanService) Run(req *model.UpdateMeetPlanRequest, resp *model.UpdateMeetPlanResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
