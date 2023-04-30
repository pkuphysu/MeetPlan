package service

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

func (h *DeleteMeetPlansService) Run(req *model.DeleteMeetPlansRequest, resp *model.DeleteMeetPlansResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
