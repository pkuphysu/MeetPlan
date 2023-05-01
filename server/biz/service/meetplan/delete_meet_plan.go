package meetplan

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

// Run req should not be nil and resp should not be nil
func (h *DeleteMeetPlanService) Run(req *model.DeleteMeetPlanRequest, resp *model.DeleteMeetPlanResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.DeleteMeetPlanResponse)
	}
	// todo edit your code
	return
}
