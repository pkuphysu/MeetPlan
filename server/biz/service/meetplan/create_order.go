package meetplan

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/biz/gorm_gen"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type CreateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IOrderDo
}

func NewCreateOrderService(ctx context.Context, RequestContext *app.RequestContext) *CreateOrderService {
	return &CreateOrderService{RequestContext: RequestContext, Context: ctx, DAO: query.Order.WithContext(ctx)}
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
	e := h.DAO.Create(&gorm_gen.Order{
		PlanID:    req.MeetPlanId,
		StudentID: req.StudentId,
		Message:   &req.Message,
	})
	if e != nil {
		return
	}
	return
}
