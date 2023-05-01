package meetplan

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

// Run req should not be nil and resp should not be nil
func (h *UpdateOrderService) Run(req *model.UpdateOrderRequest, resp *model.UpdateOrderResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.UpdateOrderResponse)
	}
	// todo edit your code
	return
}
