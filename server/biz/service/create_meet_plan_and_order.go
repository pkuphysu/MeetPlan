package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type CreateMeetPlanAndOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateMeetPlanAndOrderService(ctx context.Context, RequestContext *app.RequestContext) *CreateMeetPlanAndOrderService {
	return &CreateMeetPlanAndOrderService{RequestContext: RequestContext, Context: ctx}
}

// Run req should not be nil and resp should not be nil
func (h *CreateMeetPlanAndOrderService) Run(req *model.CreateMeetPlanAndOrderRequest, resp *model.CreateMeetPlanAndOrderResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.CreateMeetPlanAndOrderResponse)
	}
	// todo edit your code
	return
}
