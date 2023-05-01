package meetplan

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

// Run req should not be nil and resp should not be nil
func (h *CreateMeetPlanService) Run(req *model.CreateMeetPlanRequest, resp *model.CreateMeetPlanResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.CreateMeetPlanResponse)
	}
	// todo edit your code
	return
}
