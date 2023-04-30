package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	model "meetplan/biz/model"
	"meetplan/biz/service"
	"meetplan/pkg/httputil"
)

// GetMeetPlan .
// @router /api/v1/meetplan/:id [GET]
func GetMeetPlan(ctx context.Context, c *app.RequestContext) {
	var req model.GetMeetPlanRequest
	resp := new(model.GetMeetPlanResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}
	err := service.NewGetMeetPlanService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// ListMeetPlan .
// @router /api/v1/meetplan [GET]
func ListMeetPlan(ctx context.Context, c *app.RequestContext) {
	var req model.ListMeetPlanRequest
	resp := new(model.ListMeetPlanResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}
	err := service.NewListMeetPlanService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// CreateMeetPlan .
// @router /api/v1/meetplan [POST]
func CreateMeetPlan(ctx context.Context, c *app.RequestContext) {
	var req model.CreateMeetPlanRequest
	resp := new(model.CreateMeetPlanResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}
	err := service.NewCreateMeetPlanService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateMeetPlan .
// @router /api/v1/meetplan/:id [PUT]
func UpdateMeetPlan(ctx context.Context, c *app.RequestContext) {
	var req model.UpdateMeetPlanRequest
	resp := new(model.UpdateMeetPlanResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}
	err := service.NewUpdateMeetPlanService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteMeetPlan .
// @router /api/v1/meetplan/:id [DELETE]
func DeleteMeetPlan(ctx context.Context, c *app.RequestContext) {
	var req model.DeleteMeetPlanRequest
	resp := new(model.DeleteMeetPlanResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}
	err := service.NewDeleteMeetPlanService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteMeetPlans .
// @router /api/v1/meetplan [DELETE]
func DeleteMeetPlans(ctx context.Context, c *app.RequestContext) {
	var req model.DeleteMeetPlansRequest
	resp := new(model.DeleteMeetPlansResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}
	err := service.NewDeleteMeetPlansService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// GetOrder .
// @router /api/v1/order/:id [GET]
func GetOrder(ctx context.Context, c *app.RequestContext) {
	var req model.GetOrderRequest
	resp := new(model.GetOrderResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}
	err := service.NewGetOrderService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// ListOrder .
// @router /api/v1/order [GET]
func ListOrder(ctx context.Context, c *app.RequestContext) {
	var req model.ListOrderRequest
	resp := new(model.ListOrderResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}
	err := service.NewListOrderService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// CreateOrder .
// @router /api/v1/order [POST]
func CreateOrder(ctx context.Context, c *app.RequestContext) {
	var req model.CreateOrderRequest
	resp := new(model.CreateOrderResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}
	err := service.NewCreateOrderService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateOrder .
// @router /api/v1/order/:id [PUT]
func UpdateOrder(ctx context.Context, c *app.RequestContext) {
	var req model.UpdateOrderRequest
	resp := new(model.UpdateOrderResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}
	err := service.NewUpdateOrderService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// CreateMeetPlanAndOrder .
// @router /api/v1/meetplanorder [POST]
func CreateMeetPlanAndOrder(ctx context.Context, c *app.RequestContext) {
	var req model.CreateMeetPlanAndOrderRequest
	resp := new(model.CreateMeetPlanAndOrderResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}
	err := service.NewCreateMeetPlanAndOrderService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}
