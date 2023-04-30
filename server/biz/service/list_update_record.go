package service

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

func (h *ListUpdateRecordService) Run(req *model.ListUpdateRecordRequest, resp *model.ListUpdateRecordResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
