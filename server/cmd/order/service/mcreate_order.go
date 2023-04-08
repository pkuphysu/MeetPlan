package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/model"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/order"
)

type MCreateOrderServiceI interface {
	MCreateOrder(req *order.MCreateOrderReq) ([]*order.Order, error)
}

func NewMCreateOrderService(ctx context.Context) MCreateOrderServiceI {
	return &mCreateOrderService{
		ctx: ctx,
		dao: query.Q.WithContext(ctx).Order,
	}
}

type mCreateOrderService struct {
	dao query.IOrderDo
	ctx context.Context
}

func (s *mCreateOrderService) MCreateOrder(req *order.MCreateOrderReq) ([]*order.Order, error) {
	var orders []*model.Order
	for _, ro := range req.Orders {
		o := model.Order{
			PlanID:    ro.PlanId,
			StudentID: ro.StudentId,
			Message:   ro.Message,
			Status:    int8(order.OrderStatus_CREATED),
		}
		if ro.Status != nil {
			o.Status = int8(*ro.Status)
		}
		orders = append(orders, &o)
	}

	err := s.dao.CreateInBatches(orders, 1000)
	if err != nil {
		return nil, err
	}
	return packOrders(orders), nil
}
