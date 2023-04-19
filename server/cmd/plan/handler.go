package main

import (
	"context"
	"meetplan/cmd/plan/service"
	"meetplan/kitex_gen/pkuphy/meetplan/plan"
	"meetplan/pkg/errno"
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

// GetOrder implements the ServiceImpl interface.
func (s *ServiceImpl) GetOrder(ctx context.Context, req *plan.GetOrderReq) (resp *plan.GetOrderResp, err error) {
	resp = plan.NewGetOrderResp()
	if err := req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	o, err := service.NewGetOrderService(ctx).GetOrder(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Order = o
	return resp, nil
}

// MGetOrder implements the ServiceImpl interface.
func (s *ServiceImpl) MGetOrder(ctx context.Context, req *plan.MGetOrderReq) (resp *plan.MGetOrderResp, err error) {
	resp = plan.NewMGetOrderResp()
	if err := req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	orders, err := service.NewMGetOrderService(ctx).MGetOrder(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Orders = orders
	return resp, nil
}

// QueryOrder implements the ServiceImpl interface.
func (s *ServiceImpl) QueryOrder(ctx context.Context, req *plan.QueryOrderReq) (resp *plan.QueryOrderResp, err error) {
	resp = plan.NewQueryOrderResp()
	if err := req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	orders, pageParam, err := service.NewQueryOrderService(ctx).QueryOrder(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Orders = orders
	resp.PageParam = pageParam
	return resp, nil
}

// CreateOrder implements the ServiceImpl interface.
func (s *ServiceImpl) CreateOrder(ctx context.Context, req *plan.CreateOrderReq) (resp *plan.CreateOrderResp, err error) {
	resp = plan.NewCreateOrderResp()
	if err := req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	o, err := service.NewCreateOrderService(ctx).CreateOrder(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Order = o
	return resp, nil
}

// MCreateOrder implements the ServiceImpl interface.
func (s *ServiceImpl) MCreateOrder(ctx context.Context, req *plan.MCreateOrderReq) (resp *plan.MCreateOrderResp, err error) {
	resp = plan.NewMCreateOrderResp()
	if err := req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	orders, err := service.NewMCreateOrderService(ctx).MCreateOrder(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Orders = orders
	return resp, nil
}

// UpdateOrder implements the ServiceImpl interface.
func (s *ServiceImpl) UpdateOrder(ctx context.Context, req *plan.UpdateOrderReq) (resp *plan.UpdateOrderResp, err error) {
	resp = plan.NewUpdateOrderResp()
	if err := req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	o, err := service.NewUpdateOrderService(ctx).UpdateOrder(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Order = o
	return resp, nil
}
