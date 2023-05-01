package meetplan

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

// Run req should not be nil and resp should not be nil
func (h *CreateOrderService) Run(req *model.CreateOrderRequest, resp *model.CreateOrderResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateOrderService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.CreateOrderResponse)
	}
	// todo edit your code
	return
}
