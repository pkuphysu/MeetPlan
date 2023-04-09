package service

import (
	"github.com/pkuphysu/meetplan/gorm_gen/model"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/plan"
	"github.com/samber/lo"
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
		Id:             &p.ID,
		TeacherId:      &p.TeacherID,
		StartTime:      lo.ToPtr(p.StartTime.Unix()),
		Duration:       &p.Duration,
		Place:          &p.Place,
		Message:        p.Message,
		Quota:          &p.Quota,
		RemainingQuota: nil,
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
		Id:             &p.ID,
		TeacherId:      &p.TeacherID,
		StartTime:      lo.ToPtr(p.StartTime.Unix()),
		Duration:       &p.Duration,
		Place:          &p.Place,
		Message:        p.Message,
		Quota:          &p.Quota,
		RemainingQuota: lo.ToPtr(p.QuotaLeft),
	}
}
