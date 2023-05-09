package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	model "meetplan/biz/model"
	"meetplan/biz/service/option"
	"meetplan/pkg/httputil"
)

// ListFriendLink .
// @router /api/v1/friendlink [GET]
func ListFriendLink(ctx context.Context, c *app.RequestContext) {
	var req model.ListFriendLinkRequest
	resp := new(model.ListFriendLinkResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := option.NewListFriendLinkService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// CreateFriendLink .
// @router /api/v1/friendlink [POST]
func CreateFriendLink(ctx context.Context, c *app.RequestContext) {
	var req model.CreateFriendLinkRequest
	resp := new(model.CreateFriendLinkResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := option.NewCreateFriendLinkService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// ListUpdateRecord .
// @router /api/v1/updaterecord [GET]
func ListUpdateRecord(ctx context.Context, c *app.RequestContext) {
	var req model.ListUpdateRecordRequest
	resp := new(model.ListUpdateRecordResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := option.NewListUpdateRecordService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// CreateUpdateRecord .
// @router /api/v1/updaterecord [POST]
func CreateUpdateRecord(ctx context.Context, c *app.RequestContext) {
	var req model.CreateUpdateRecordRequest
	resp := new(model.CreateUpdateRecordResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := option.NewCreateUpdateRecordService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// GetTermDateRange .
// @router /api/v1/termdate [GET]
func GetTermDateRange(ctx context.Context, c *app.RequestContext) {
	var req model.GetTermDateRangeRequest
	resp := new(model.GetTermDateRangeResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := option.NewGetTermDateRangeService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateTermDateRange .
// @router /api/v1/termdate [PUT]
func UpdateTermDateRange(ctx context.Context, c *app.RequestContext) {
	var req model.UpdateTermDateRangeRequest
	resp := new(model.UpdateTermDateRangeResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := option.NewUpdateTermDateRangeService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// GetOption .
// @router /api/v1/option [GET]
func GetOption(ctx context.Context, c *app.RequestContext) {
	var req model.GetOptionRequest
	resp := new(model.GetOptionResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := option.NewGetOptionService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateOption .
// @router /api/v1/option [PUT]
func UpdateOption(ctx context.Context, c *app.RequestContext) {
	var req model.UpdateOptionRequest
	resp := new(model.UpdateOptionResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := option.NewUpdateOptionService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}