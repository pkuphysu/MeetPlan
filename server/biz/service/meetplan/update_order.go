package meetplan

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gen/field"
	"gorm.io/gorm"

	"meetplan/biz/dal/pack"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type UpdateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IOrderDo
}

func NewUpdateOrderService(ctx context.Context, RequestContext *app.RequestContext) *UpdateOrderService {
	return &UpdateOrderService{RequestContext: RequestContext, Context: ctx, DAO: query.Order.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *UpdateOrderService) Run(req *model.UpdateOrderRequest, resp *model.UpdateOrderResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in UpdateOrderService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.UpdateOrderResponse)
	}

	var fields []field.AssignExpr
	if req.Message != nil {
		fields = append(fields, query.Order.Message.Value(*req.Message))
	}
	if req.Status != nil {
		fields = append(fields, query.Order.Status.Value(int8(*req.Status)))
	}
	if len(fields) == 0 {
		return errno.NewValidationErr("no field to update")
	}

	_, e := h.DAO.Where(query.Order.ID.Eq(req.Id)).UpdateColumnSimple(fields...)
	if e != nil && errors.Is(e, gorm.ErrRecordNotFound) {
		return errno.NewNotFoundErr("order not found")
	} else if e != nil {
		return errno.ToInternalErr(e)
	}

	order, e := h.DAO.Preload(query.Order.Plan, query.Order.Student).Where(query.Order.ID.Eq(req.Id)).First()
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = pack.OrderDal2Vo(order)
	return
}
