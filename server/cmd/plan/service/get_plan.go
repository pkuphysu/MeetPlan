package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/plan"
)

type GetPlanServiceI interface {
	GetPlan(req *plan.GetPlanReq) (*plan.Plan, error)
}

func NewGetPlanService(ctx context.Context) GetPlanServiceI {
	return &getPlanService{
		dao: query.Q.WithContext(ctx).Plan,
	}
}

type getPlanService struct {
	dao query.IPlanDo
}

func (s *getPlanService) GetPlan(req *plan.GetPlanReq) (*plan.Plan, error) {
	p, err := s.dao.Where(query.Q.Plan.ID.Eq(req.Id)).First()
	if err != nil {
		return nil, err
	}
	return packPlan(p), nil
}
