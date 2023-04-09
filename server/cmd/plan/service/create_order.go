package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/model"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/plan"
)

type CreateOrderServiceI interface {
	CreateOrder(req *plan.CreateOrderReq) (*plan.Order, error)
}

func NewCreateOrderService(ctx context.Context) CreateOrderServiceI {
	return &createOrderService{
		ctx: ctx,
		dao: query.Q.WithContext(ctx).Order,
	}
}

type createOrderService struct {
	dao query.IOrderDo
	ctx context.Context
}

func (s *createOrderService) CreateOrder(req *plan.CreateOrderReq) (*plan.Order, error) {
	o := model.Order{
		PlanID:    req.PlanId,
		StudentID: req.StudentId,
		Message:   req.Message,
		Status:    int8(plan.OrderStatus_CREATED),
	}
	if req.Status != nil {
		o.Status = int8(*req.Status)
	}

	err := s.dao.Create(&o)
	if err != nil {
		return nil, err
	}
	return packOrder(&o), nil
}
