package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/plan"
)

type MGetPlanServiceI interface {
	MGetPlan(req *plan.MGetPlanReq) ([]*plan.Plan, error)
}

func NewMGetPlanService(ctx context.Context) MGetPlanServiceI {
	return &mGetPlanService{
		dao: query.Q.WithContext(ctx).Plan,
	}
}

type mGetPlanService struct {
	dao query.IPlanDo
}

func (s *mGetPlanService) MGetPlan(req *plan.MGetPlanReq) ([]*plan.Plan, error) {
	plans, err := s.dao.Where(query.Q.Plan.ID.In(req.IdList...)).Find()
	if err != nil {
		return nil, err
	}
	return packPlans(plans), nil
}
