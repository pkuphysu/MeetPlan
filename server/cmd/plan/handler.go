package main

import (
	"context"
	"github.com/pkuphysu/meetplan/cmd/plan/service"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/plan"
	"github.com/pkuphysu/meetplan/pkg/errno"
)

// ServiceImpl implements the last service interface defined in the IDL.
type ServiceImpl struct{}

// GetPlan implements the ServiceImpl interface.
func (s *ServiceImpl) GetPlan(ctx context.Context, req *plan.GetPlanReq) (resp *plan.GetPlanResp, err error) {
	resp = plan.NewGetPlanResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	p, err := service.NewGetPlanService(ctx).GetPlan(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Plan = p
	return resp, nil
}

// MGetPlan implements the ServiceImpl interface.
func (s *ServiceImpl) MGetPlan(ctx context.Context, req *plan.MGetPlanReq) (resp *plan.MGetPlanResp, err error) {
	resp = plan.NewMGetPlanResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	plans, err := service.NewMGetPlanService(ctx).MGetPlan(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.PlanList = plans
	return resp, nil
}

// QueryPlan implements the ServiceImpl interface.
func (s *ServiceImpl) QueryPlan(ctx context.Context, req *plan.QueryPlanReq) (resp *plan.QueryPlanResp, err error) {
	resp = plan.NewQueryPlanResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	plans, pageParam, err := service.NewQueryPlanService(ctx).QueryPlan(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.PlanList = plans
	resp.PageParam = pageParam
	return resp, nil
}

// CreatePlan implements the ServiceImpl interface.
func (s *ServiceImpl) CreatePlan(ctx context.Context, req *plan.CreatePlanReq) (resp *plan.CreatePlanResp, err error) {
	resp = plan.NewCreatePlanResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}
	p, err := service.NewCreatePlanService(ctx).CreatePlan(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Plan = p
	return resp, nil
}

// MCreatePlan implements the ServiceImpl interface.
func (s *ServiceImpl) MCreatePlan(ctx context.Context, req *plan.MCreatePlanReq) (resp *plan.MCreatePlanResp, err error) {
	resp = plan.NewMCreatePlanResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}
	plans, err := service.NewMCreatePlanService(ctx).MCreatePlan(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.PlanList = plans
	return resp, nil
}
