package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type DeleteMeetPlanService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteMeetPlanService(ctx context.Context, RequestContext *app.RequestContext) *DeleteMeetPlanService {
	return &DeleteMeetPlanService{RequestContext: RequestContext, Context: ctx}
}

func (h *DeleteMeetPlanService) Run(req *model.DeleteMeetPlanRequest, resp *model.DeleteMeetPlanResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
