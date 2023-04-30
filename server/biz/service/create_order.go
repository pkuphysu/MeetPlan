package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type CreateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateOrderService(ctx context.Context, RequestContext *app.RequestContext) *CreateOrderService {
	return &CreateOrderService{RequestContext: RequestContext, Context: ctx}
}

func (h *CreateOrderService) Run(req *model.CreateOrderRequest, resp *model.CreateOrderResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
