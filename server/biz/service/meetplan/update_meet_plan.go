package meetplan

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

// Run req should not be nil and resp should not be nil
func (h *UpdateMeetPlanService) Run(req *model.UpdateMeetPlanRequest, resp *model.UpdateMeetPlanResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in UpdateMeetPlanService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.UpdateMeetPlanResponse)
	}
	// todo edit your code
	return
}
