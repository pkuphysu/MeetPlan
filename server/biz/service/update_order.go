package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type UpdateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateOrderService(ctx context.Context, RequestContext *app.RequestContext) *UpdateOrderService {
	return &UpdateOrderService{RequestContext: RequestContext, Context: ctx}
}

func (h *UpdateOrderService) Run(req *model.UpdateOrderRequest, resp *model.UpdateOrderResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
