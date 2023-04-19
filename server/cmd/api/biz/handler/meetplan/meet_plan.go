package meetplan

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	model "meetplan/cmd/api/biz/model"
	"meetplan/cmd/api/utils/httputil"
)

// GetMeetPlan .
// @router /api/v1/meetplan/:id [GET]
func GetMeetPlan(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.GetMeetPlanRequest
	resp := new(model.GetMeetPlanResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	// add your logic here
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// ListMeetPlan .
// @router /api/v1/meetplan [GET]
func ListMeetPlan(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.ListMeetPlanRequest
	resp := new(model.ListMeetPlanResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	// add your logic here
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// CreateMeetPlan .
// @router /api/v1/meetplan [POST]
func CreateMeetPlan(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.CreateMeetPlanRequest
	resp := new(model.CreateMeetPlanResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	// add your logic here
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateMeetPlan .
// @router /api/v1/meetplan/:id [PUT]
func UpdateMeetPlan(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.UpdateMeetPlanRequest
	resp := new(model.UpdateMeetPlanResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	// add your logic here
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteMeetPlan .
// @router /api/v1/meetplan/:id [DELETE]
func DeleteMeetPlan(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.DeleteMeetPlanRequest
	resp := new(model.DeleteMeetPlanResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	// add your logic here
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteMeetPlans .
// @router /api/v1/meetplan [DELETE]
func DeleteMeetPlans(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.DeleteMeetPlansRequest
	resp := new(model.DeleteMeetPlansResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	// add your logic here
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// GetOrder .
// @router /api/v1/order/:id [GET]
func GetOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.GetOrderRequest
	resp := new(model.GetOrderResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	// add your logic here
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// ListOrder .
// @router /api/v1/order [GET]
func ListOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.ListOrderRequest
	resp := new(model.ListOrderResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	// add your logic here
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// CreateOrder .
// @router /api/v1/order [POST]
func CreateOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.CreateOrderRequest
	resp := new(model.CreateOrderResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	// add your logic here
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateOrder .
// @router /api/v1/order/:id [PUT]
func UpdateOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.UpdateOrderRequest
	resp := new(model.UpdateOrderResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	// add your logic here
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// CreateMeetPlanAndOrder .
// @router /api/v1/meetplanorder [POST]
func CreateMeetPlanAndOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.CreateMeetPlanAndOrderRequest
	resp := new(model.CreateMeetPlanAndOrderResponse)
	err = c.BindAndValidate(&req)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	// add your logic here
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}
