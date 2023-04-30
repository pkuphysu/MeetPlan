package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type CreateMeetPlanService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateMeetPlanService(ctx context.Context, RequestContext *app.RequestContext) *CreateMeetPlanService {
	return &CreateMeetPlanService{RequestContext: RequestContext, Context: ctx}
}

func (h *CreateMeetPlanService) Run(req *model.CreateMeetPlanRequest, resp *model.CreateMeetPlanResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
