package service

import (
	"context"
	"meetplan/cmd/plan/constants"
	"meetplan/cmd/plan/rpc"
	"meetplan/gorm_gen/model"
	"meetplan/gorm_gen/query"
	"meetplan/kitex_gen/pkuphy/meetplan/plan"
	"meetplan/pkg/errno"
	"time"
)

type CreatePlanServiceI interface {
	CreatePlan(req *plan.CreatePlanReq) (*plan.Plan, error)
}

func NewCreatePlanService(ctx context.Context) CreatePlanServiceI {
	return &createPlanService{
		ctx: ctx,
		dao: query.Q.WithContext(ctx).Plan,
	}
}

type createPlanService struct {
	ctx context.Context
	dao query.IPlanDo
}

func (s *createPlanService) CreatePlan(req *plan.CreatePlanReq) (*plan.Plan, error) {
	p := model.Plan{}
	if req.Plan.TeacherId == nil || req.Plan.StartTime == nil {
		return nil, errno.ParamErr.WithMessage("teacher_id and start_time are required")
	}
	p.TeacherID = *req.Plan.TeacherId
	p.StartTime = time.Unix(*req.Plan.StartTime, 0)

	tea, err := rpc.GetUserDetail(s.ctx, p.TeacherID)
	if err != nil {
		return nil, err
	}
	if tea.IsTeacher == nil || !*tea.IsTeacher {
		return nil, errno.ParamErr.WithMessage("teacher_id is not a teacher")
	}

	if req.Plan.Duration != nil {
		p.Duration = *req.Plan.Duration
	} else {
		p.Duration = 30 * 60
	}
	if req.Plan.Place != nil {
		p.Place = *req.Plan.Place
	} else if tea.Office != nil {
		p.Place = *tea.Office
	} else {
		return nil, errno.ParamErr.WithMessage("place is required")
	}

	if req.Plan.Message != nil {
		p.Message = req.Plan.Message
	}
	if req.Plan.Quota != nil {
		p.Quota = *req.Plan.Quota
	} else {
		p.Quota = constants.DefaultQuota
	}

	err = s.dao.Create(&p)
	if err != nil {
		return nil, err
	}
	return packPlan(&p), nil
}
