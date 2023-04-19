package service

import (
	"context"
	"meetplan/gorm_gen/query"
	"meetplan/kitex_gen/pkuphy/meetplan/plan"
)

type MGetPlanServiceI interface {
	MGetPlan(req *plan.MGetPlanReq) ([]*plan.Plan, error)
}

func NewMGetPlanService(ctx context.Context) MGetPlanServiceI {
	return &mGetPlanService{
		dao: query.Q.WithContext(ctx).PlanView,
	}
}

type mGetPlanService struct {
	dao query.IPlanViewDo
}

func (s *mGetPlanService) MGetPlan(req *plan.MGetPlanReq) ([]*plan.Plan, error) {
	plans, err := s.dao.Where(query.Q.PlanView.ID.In(req.IdList...)).Find()
	if err != nil {
		return nil, err
	}
	return packPlanViews(plans), nil
}
