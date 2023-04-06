package main

import (
	"context"
	order "github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/order"
)

// ServiceImpl implements the last service interface defined in the IDL.
type ServiceImpl struct{}

// GetOrder implements the ServiceImpl interface.
func (s *ServiceImpl) GetOrder(ctx context.Context, req *order.GetOrderReq) (resp *order.GetOrderResp, err error) {
	// TODO: Your code here...
	return
}

// MGetOrder implements the ServiceImpl interface.
func (s *ServiceImpl) MGetOrder(ctx context.Context, req *order.MGetOrderReq) (resp *order.MGetOrderResp, err error) {
	// TODO: Your code here...
	return
}

// QueryOrder implements the ServiceImpl interface.
func (s *ServiceImpl) QueryOrder(ctx context.Context, req *order.QueryOrderReq) (resp *order.QueryOrderResp, err error) {
	// TODO: Your code here...
	return
}

// CreateOrder implements the ServiceImpl interface.
func (s *ServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (resp *order.CreateOrderResp, err error) {
	// TODO: Your code here...
	return
}

// MCreateOrder implements the ServiceImpl interface.
func (s *ServiceImpl) MCreateOrder(ctx context.Context, req *order.MCreateOrderReq) (resp *order.MCreateOrderResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateOrder implements the ServiceImpl interface.
func (s *ServiceImpl) UpdateOrder(ctx context.Context, req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	// TODO: Your code here...
	return
}
