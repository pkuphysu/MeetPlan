package meetplan

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

// Run req should not be nil and resp should not be nil
func (h *GetOrderService) Run(req *model.GetOrderRequest, resp *model.GetOrderResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.GetOrderResponse)
	}
	// todo edit your code
	return
}
