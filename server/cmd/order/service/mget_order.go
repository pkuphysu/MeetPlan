package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/order"
)

type MGetOrderServiceI interface {
	MGetOrder(req *order.MGetOrderReq) ([]*order.Order, error)
}

func NewMGetOrderService(ctx context.Context) MGetOrderServiceI {
	return &mGetOrderService{
		ctx: ctx,
		dao: query.Q.WithContext(ctx).Order,
	}
}

type mGetOrderService struct {
	dao query.IOrderDo
	ctx context.Context
}

func (s *mGetOrderService) MGetOrder(req *order.MGetOrderReq) ([]*order.Order, error) {
	orders, err := s.dao.Where(query.Q.Order.ID.In(req.Ids...)).Find()
	if err != nil {
		return nil, err
	}
	return packOrders(orders), nil
}
