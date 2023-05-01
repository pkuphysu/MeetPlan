package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	model "meetplan/biz/model"
	"meetplan/biz/service/user"
	"meetplan/pkg/httputil"
)

// Login .
// @router /api/v1/login [GET]
func Login(ctx context.Context, c *app.RequestContext) {
	var req model.LoginRequest
	resp := new(model.LoginResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := user.NewLoginService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// GetSelf .
// @router /api/v1/user/self [GET]
func GetSelf(ctx context.Context, c *app.RequestContext) {
	var req model.GetSelfRequest
	resp := new(model.GetSelfResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := user.NewGetSelfService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// GetUser .
// @router /api/v1/user/:id [GET]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var req model.GetUserRequest
	resp := new(model.GetUserResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := user.NewGetUserService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// ListUser .
// @router /api/v1/user [GET]
func ListUser(ctx context.Context, c *app.RequestContext) {
	var req model.ListUserRequest
	resp := new(model.ListUserResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := user.NewListUserService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}

// CreateUser .
// @router /api/v1/user [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var req model.CreateUserRequest
	resp := new(model.CreateUserResponse)
	if err := c.BindAndValidate(&req); err != nil {
		resp.Code = -1
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, consts.StatusBadRequest, resp)
		return
	}

	err := user.NewCreateUserService(ctx, c).Run(&req, resp)

	if err != nil {
		resp.Code = int32(err.ErrCode())
		resp.Message = err.Error()
		httputil.SendResponse(ctx, c, err.StatusCode(), resp)
		return
	}
	httputil.SendResponse(ctx, c, consts.StatusOK, resp)
}
