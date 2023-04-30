package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type GetOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetOrderService(ctx context.Context, RequestContext *app.RequestContext) *GetOrderService {
	return &GetOrderService{RequestContext: RequestContext, Context: ctx}
}

func (h *GetOrderService) Run(req *model.GetOrderRequest, resp *model.GetOrderResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
