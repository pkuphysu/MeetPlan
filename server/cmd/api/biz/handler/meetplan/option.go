package meetplan

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	model "meetplan/cmd/api/biz/model"
	"meetplan/cmd/api/utils/httputil"
)

// ListFriendLink .
// @router /api/v1/friendlink [GET]
func ListFriendLink(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.ListFriendLinkRequest
	resp := new(model.ListFriendLinkResponse)
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

// CreateFriendLink .
// @router /api/v1/friendlink [POST]
func CreateFriendLink(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.CreateFriendLinkRequest
	resp := new(model.CreateFriendLinkResponse)
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

// ListUpdateRecord .
// @router /api/v1/updaterecord [GET]
func ListUpdateRecord(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.ListUpdateRecordRequest
	resp := new(model.ListUpdateRecordResponse)
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

// CreateUpdateRecord .
// @router /api/v1/updaterecord [POST]
func CreateUpdateRecord(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.CreateUpdateRecordRequest
	resp := new(model.CreateUpdateRecordResponse)
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

// GetTermDateRange .
// @router /api/v1/termdate [GET]
func GetTermDateRange(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.GetTermDateRangeRequest
	resp := new(model.GetTermDateRangeResponse)
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

// UpdateTermDateRange .
// @router /api/v1/termdate [PUT]
func UpdateTermDateRange(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.UpdateTermDateRangeRequest
	resp := new(model.UpdateTermDateRangeResponse)
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
