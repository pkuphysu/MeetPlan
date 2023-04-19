package service

import (
	"github.com/samber/lo"
	"meetplan/gorm_gen/model"
	"meetplan/kitex_gen/pkuphy/meetplan/plan"
)

func packPlans(plans []*model.Plan) []*plan.Plan {
	var resp []*plan.Plan
	for _, p := range plans {
		resp = append(resp, packPlan(p))
	}
	return resp
}

func packPlan(p *model.Plan) *plan.Plan {
	return &plan.Plan{
		Id:        &p.ID,
		TeacherId: &p.TeacherID,
		StartTime: lo.ToPtr(p.StartTime.Unix()),
		Duration:  &p.Duration,
		Place:     &p.Place,
		Message:   p.Message,
		Quota:     &p.Quota,
	}
}

func packPlanViews(plans []*model.PlanView) []*plan.Plan {
	var resp []*plan.Plan
	for _, p := range plans {
		resp = append(resp, packPlanView(p))
	}
	return resp
}

func packPlanView(p *model.PlanView) *plan.Plan {
	return &plan.Plan{
		Id:        &p.ID,
		TeacherId: &p.TeacherID,
		StartTime: lo.ToPtr(p.StartTime.Unix()),
		Duration:  &p.Duration,
		Place:     &p.Place,
		Message:   p.Message,
		Quota:     &p.Quota,
	}
}
func packOrders(orders []*model.Order) []*plan.Order {
	var resp []*plan.Order
	for _, o := range orders {
		resp = append(resp, packOrder(o))
	}
	return resp
}

func packOrder(o *model.Order) *plan.Order {
	return &plan.Order{
		Id:        &o.ID,
		PlanId:    &o.PlanID,
		StudentId: &o.StudentID,
		Message:   o.Message,
		Status:    lo.ToPtr(plan.OrderStatus(o.Status)),
	}
}
