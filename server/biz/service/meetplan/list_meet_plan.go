package meetplan

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

// Run req should not be nil and resp should not be nil
func (h *ListMeetPlanService) Run(req *model.ListMeetPlanRequest, resp *model.ListMeetPlanResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.ListMeetPlanResponse)
	}
	// todo edit your code
	return
}
