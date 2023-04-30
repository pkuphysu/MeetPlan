package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type UpdateTermDateRangeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateTermDateRangeService(ctx context.Context, RequestContext *app.RequestContext) *UpdateTermDateRangeService {
	return &UpdateTermDateRangeService{RequestContext: RequestContext, Context: ctx}
}

// Run req should not be nil and resp should not be nil
func (h *UpdateTermDateRangeService) Run(req *model.UpdateTermDateRangeRequest, resp *model.UpdateTermDateRangeResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.UpdateTermDateRangeResponse)
	}
	// todo edit your code
	return
}
