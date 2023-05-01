package option

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type ListUpdateRecordService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListUpdateRecordService(ctx context.Context, RequestContext *app.RequestContext) *ListUpdateRecordService {
	return &ListUpdateRecordService{RequestContext: RequestContext, Context: ctx}
}

// Run req should not be nil and resp should not be nil
func (h *ListUpdateRecordService) Run(req *model.ListUpdateRecordRequest, resp *model.ListUpdateRecordResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.ListUpdateRecordResponse)
	}
	// todo edit your code
	return
}
