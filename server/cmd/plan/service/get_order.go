package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/plan"
)

type GetOrderServiceI interface {
	GetOrder(req *plan.GetOrderReq) (*plan.Order, error)
}

func NewGetOrderService(ctx context.Context) GetOrderServiceI {
	return &getOrderService{
		ctx: ctx,
		dao: query.Q.WithContext(ctx).Order,
	}
}

type getOrderService struct {
	dao query.IOrderDo
	ctx context.Context
}

func (s *getOrderService) GetOrder(req *plan.GetOrderReq) (*plan.Order, error) {
	o, err := s.dao.Where(query.Q.Order.ID.Eq(req.Id)).First()
	if err != nil {
		return nil, err
	}
	return packOrder(o), nil
}
