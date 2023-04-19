package service

import (
	"context"
	"meetplan/gorm_gen/query"
	"meetplan/kitex_gen/pkuphy/meetplan/plan"
	"meetplan/pkg/errno"
)

type UpdateOrderServiceI interface {
	UpdateOrder(req *plan.UpdateOrderReq) (*plan.Order, error)
}

func NewUpdateOrderService(ctx context.Context) UpdateOrderServiceI {
	return &updateOrderService{
		ctx: ctx,
		dao: query.Q.WithContext(ctx).Order,
	}
}

type updateOrderService struct {
	dao query.IOrderDo
	ctx context.Context
}

func (s *updateOrderService) UpdateOrder(req *plan.UpdateOrderReq) (*plan.Order, error) {
	updateMap := make(map[string]interface{})
	if req.Message != nil {
		updateMap["message"] = *req.Message
	}
	if req.Status != nil {
		updateMap["status"] = *req.Status
	}
	if len(updateMap) == 0 {
		return nil, nil
	}
	res, err := s.dao.Where(query.Q.Order.ID.Eq(req.Id)).Updates(updateMap)
	if err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errno.OrderNotFoundErr
	}
	o, err := s.dao.Where(query.Q.Order.ID.Eq(req.Id)).First()
	if err != nil {
		return nil, err
	}
	return packOrder(o), nil
}
