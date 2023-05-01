package meetplan

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

// Run req should not be nil and resp should not be nil
func (h *GetMeetPlanService) Run(req *model.GetMeetPlanRequest, resp *model.GetMeetPlanResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in GetMeetPlanService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.GetMeetPlanResponse)
	}
	// todo edit your code
	return
}
