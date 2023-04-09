package service

import (
	"context"
	"github.com/pkuphysu/meetplan/cmd/plan/constants"
	"github.com/pkuphysu/meetplan/cmd/plan/rpc"
	"github.com/pkuphysu/meetplan/gorm_gen/model"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/plan"
	"github.com/pkuphysu/meetplan/pkg/errno"
	"github.com/samber/lo"
	"time"
)

type MCreatePlanServiceI interface {
	MCreatePlan(req *plan.MCreatePlanReq) ([]*plan.Plan, error)
}

func NewMCreatePlanService(ctx context.Context) MCreatePlanServiceI {
	return &mCreatePlanService{
		ctx: ctx,
		dao: query.Q.WithContext(ctx).Plan,
	}
}

type mCreatePlanService struct {
	ctx context.Context
	dao query.IPlanDo
}

func (s *mCreatePlanService) MCreatePlan(req *plan.MCreatePlanReq) ([]*plan.Plan, error) {
	var plans []*model.Plan

	var teacherIDs []int64
	for _, p := range req.PlanList {
		if p.TeacherId == nil || p.StartTime == nil {
			return nil, errno.ParamErr.WithMessage("teacher_id and start_time are required")
		}
		teacherIDs = append(teacherIDs, *p.TeacherId)
	}
	teacherIDs = lo.Uniq(teacherIDs)
	userMap, err := rpc.MGetUserDetailMap(s.ctx, teacherIDs)
	if err != nil {
		return nil, err
	}

	for _, tmp := range req.PlanList {
		p := &model.Plan{}
		p.TeacherID = *tmp.TeacherId
		p.StartTime = time.Unix(*tmp.StartTime, 0)

		tea, ok := userMap[p.TeacherID]
		if !ok || tea.IsTeacher == nil || !*tea.IsTeacher {
			return nil, errno.ParamErr.WithMessage("teacher_id is not a teacher")
		}

		if tmp.Duration != nil {
			p.Duration = *tmp.Duration
		} else {
			p.Duration = 30 * 60
		}
		if tmp.Place != nil {
			p.Place = *tmp.Place
		} else if tea.Office != nil {
			p.Place = *tea.Office
		} else {
			return nil, errno.ParamErr.WithMessage("place is required")
		}

		if tmp.Message != nil {
			p.Message = tmp.Message
		}
		if tmp.Quota != nil {
			p.Quota = *tmp.Quota
		} else {
			p.Quota = constants.DefaultQuota
		}

		plans = append(plans, p)
	}

	err = s.dao.CreateInBatches(plans, 100)
	if err != nil {
		return nil, err
	}
	return packPlans(plans), nil
}
