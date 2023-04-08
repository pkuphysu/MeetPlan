package service

import (
	"github.com/pkuphysu/meetplan/gorm_gen/model"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/order"
	"github.com/samber/lo"
)

func packOrders(orders []*model.Order) []*order.Order {
	var resp []*order.Order
	for _, o := range orders {
		resp = append(resp, packOrder(o))
	}
	return resp
}

func packOrder(o *model.Order) *order.Order {
	return &order.Order{
		Id:        &o.ID,
		PlanId:    &o.PlanID,
		StudentId: &o.StudentID,
		Message:   o.Message,
		Status:    lo.ToPtr(order.OrderStatus(o.Status)),
	}
}
