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

type GetOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	OrderDAO       query.IOrderDo
}

func NewGetOrderService(ctx context.Context, RequestContext *app.RequestContext) *GetOrderService {
	return &GetOrderService{RequestContext: RequestContext, Context: ctx, OrderDAO: query.Order.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *GetOrderService) Run(req *model.GetOrderRequest, resp *model.GetOrderResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in GetOrderService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.GetOrderResponse)
	}
	var preloads []field.RelationField
	if req.WithTeacher {
		preloads = append(preloads, query.Order.Teacher)
	}
	if req.WithStudent {
		preloads = append(preloads, query.Order.Student)
	}
	if req.WithMeetPlan {
		preloads = append(preloads, query.Order.Plan)
	}
	order, e := h.OrderDAO.Preload(preloads...).Where(query.Order.ID.Eq(req.Id)).First()
	if e != nil && errors.Is(e, gorm.ErrRecordNotFound) {
		return errno.NewNotFoundErr("plan order not found")
	} else if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = pack.OrderDal2Vo(order)
	return
}
