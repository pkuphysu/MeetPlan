package main

import (
	"context"
	"github.com/pkuphysu/meetplan/cmd/order/service"
	order "github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/order"
	"github.com/pkuphysu/meetplan/pkg/errno"
)

// ServiceImpl implements the last service interface defined in the IDL.
type ServiceImpl struct{}

// GetOrder implements the ServiceImpl interface.
func (s *ServiceImpl) GetOrder(ctx context.Context, req *order.GetOrderReq) (resp *order.GetOrderResp, err error) {
	resp = order.NewGetOrderResp()
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
func (s *ServiceImpl) MGetOrder(ctx context.Context, req *order.MGetOrderReq) (resp *order.MGetOrderResp, err error) {
	resp = order.NewMGetOrderResp()
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
func (s *ServiceImpl) QueryOrder(ctx context.Context, req *order.QueryOrderReq) (resp *order.QueryOrderResp, err error) {
	resp = order.NewQueryOrderResp()
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
func (s *ServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (resp *order.CreateOrderResp, err error) {
	resp = order.NewCreateOrderResp()
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
func (s *ServiceImpl) MCreateOrder(ctx context.Context, req *order.MCreateOrderReq) (resp *order.MCreateOrderResp, err error) {
	resp = order.NewMCreateOrderResp()
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
func (s *ServiceImpl) UpdateOrder(ctx context.Context, req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	resp = order.NewUpdateOrderResp()
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
