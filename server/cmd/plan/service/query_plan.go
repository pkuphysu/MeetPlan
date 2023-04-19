package service

import (
	"context"
	"meetplan/gorm_gen/query"
	"meetplan/kitex_gen/pkuphy/meetplan/base"
	"meetplan/kitex_gen/pkuphy/meetplan/plan"
	"time"
)

type QueryPlanServiceI interface {
	QueryPlan(req *plan.QueryPlanReq) ([]*plan.Plan, *base.PageParam, error)
}

func NewQueryPlanService(ctx context.Context) QueryPlanServiceI {
	return &queryPlanService{
		dao: query.Q.WithContext(ctx).PlanView,
	}
}

type queryPlanService struct {
	dao query.IPlanViewDo
}

func (s *queryPlanService) QueryPlan(req *plan.QueryPlanReq) ([]*plan.Plan, *base.PageParam, error) {
	dao := s.dao
	if len(req.TeacherIdList) > 0 {
		dao = dao.Where(query.Q.PlanView.TeacherID.In(req.TeacherIdList...))
	}
	if req.StartTime != nil {
		dao = dao.Where(query.Q.PlanView.StartTime.Gte(time.Unix(*req.StartTime, 0)))
	}

	var pageParam *base.PageParam
	if req.PageParam != nil {
		pageParam = req.PageParam
	} else {
		pageParam = base.NewPageParam()
		pageParam.SetPageNum(1)
		pageParam.SetPageSize(10)
	}

	dao = dao.Limit(int(pageParam.PageSize)).Offset(int(pageParam.PageSize * (pageParam.PageNum - 1)))

	plans, err := dao.Find()
	if err != nil {
		return nil, nil, err
	}
	return packPlanViews(plans), pageParam, nil
}
