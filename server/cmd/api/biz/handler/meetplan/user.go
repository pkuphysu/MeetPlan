package meetplan

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	model "meetplan/cmd/api/biz/model"
	"meetplan/cmd/api/utils/httputil"
)

// Login .
// @router /api/v1/login [GET]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.LoginRequest
	resp := new(model.LoginResponse)
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

// GetSelf .
// @router /api/v1/user/self [GET]
func GetSelf(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.GetSelfRequest
	resp := new(model.GetSelfResponse)
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

// GetUser .
// @router /api/v1/user/:id [GET]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.GetUserRequest
	resp := new(model.GetUserResponse)
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

// ListUser .
// @router /api/v1/user [GET]
func ListUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.ListUserRequest
	resp := new(model.ListUserResponse)
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
