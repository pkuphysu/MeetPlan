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

func (h *CreateUpdateRecordService) Run(req *model.CreateUpdateRecordRequest, resp *model.CreateUpdateRecordResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
