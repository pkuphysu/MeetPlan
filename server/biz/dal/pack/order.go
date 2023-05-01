package pack

import (
	"github.com/samber/lo"

	"meetplan/biz/gorm_gen"
	"meetplan/biz/model"
)

func OrderDal2Vo(order *gorm_gen.Order) *model.Order {
	if order == nil {
		return nil
	}
	return &model.Order{
		Id:         order.ID,
		MeetPlanId: order.PlanID,
		StudentId:  order.StudentID,
		Message: lo.IfF(order.Message != nil, func() string {
			return *order.Message
		}).Else(""),
		Status:   model.OrderStatus(order.Status),
		MeetPlan: PlanDal2Vo(order.Plan),
		Student:  UserDal2Vo(order.Student),
		Teacher:  UserDal2Vo(order.Teacher),
	}
}

func OrdersDal2Vo(orders []*gorm_gen.Order) []*model.Order {
	if orders == nil {
		return nil
	}
	var res []*model.Order
	for _, order := range orders {
		res = append(res, OrderDal2Vo(order))
	}
	return res
}

func OrderVo2Dal(order *model.Order) *gorm_gen.Order {
	if order == nil {
		return nil
	}
	return &gorm_gen.Order{
		ID:        order.Id,
		Status:    int8(order.Status),
		Message:   &order.Message,
		PlanID:    order.MeetPlanId,
		StudentID: order.StudentId,
		Student:   nil,
		Teacher:   nil,
		Plan:      nil,
	}
}

func OrdersVo2Dal(orders []*model.Order) []*gorm_gen.Order {
	if orders == nil {
		return nil
	}
	var res []*gorm_gen.Order
	for _, order := range orders {
		res = append(res, OrderVo2Dal(order))
	}
	return res
}
