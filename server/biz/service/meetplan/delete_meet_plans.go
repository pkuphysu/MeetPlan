package meetplan

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type DeleteMeetPlansService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteMeetPlansService(ctx context.Context, RequestContext *app.RequestContext) *DeleteMeetPlansService {
	return &DeleteMeetPlansService{RequestContext: RequestContext, Context: ctx}
}

// Run req should not be nil and resp should not be nil
func (h *DeleteMeetPlansService) Run(req *model.DeleteMeetPlansRequest, resp *model.DeleteMeetPlansResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in DeleteMeetPlansService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.DeleteMeetPlansResponse)
	}
	// todo edit your code
	return
}
