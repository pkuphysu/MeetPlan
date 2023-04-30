package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type CreateUpdateRecordService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateUpdateRecordService(ctx context.Context, RequestContext *app.RequestContext) *CreateUpdateRecordService {
	return &CreateUpdateRecordService{RequestContext: RequestContext, Context: ctx}
}

// Run req should not be nil and resp should not be nil
func (h *CreateUpdateRecordService) Run(req *model.CreateUpdateRecordRequest, resp *model.CreateUpdateRecordResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.CreateUpdateRecordResponse)
	}
	// todo edit your code
	return
}
