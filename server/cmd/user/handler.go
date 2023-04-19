package main

import (
	"context"
	"meetplan/cmd/user/service"
	"meetplan/kitex_gen/pkuphy/meetplan/user"
	"meetplan/pkg/errno"
)

// ServiceImpl implements the last service interface defined in the IDL.
type ServiceImpl struct{}

// Login implements the ServiceImpl interface.
func (s *ServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp = user.NewLoginResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	resp, err = service.NewLoginService(ctx).Login(ctx, req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	return resp, nil
}

// GetUser implements the ServiceImpl interface.
func (s *ServiceImpl) GetUser(ctx context.Context, req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	resp = user.NewGetUserResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	u, err := service.NewGetUserService(ctx).GetUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.User = u
	return resp, nil
}

// MGetUser implements the ServiceImpl interface.
func (s *ServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserReq) (resp *user.MGetUserResp, err error) {
	resp = user.NewMGetUserResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
}

// QueryUser implements the ServiceImpl interface.
func (s *ServiceImpl) QueryUser(ctx context.Context, req *user.QueryUserReq) (resp *user.QueryUserResp, err error) {
	resp = user.NewQueryUserResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	users, pageParam, err := service.NewQueryUserService(ctx).QueryUser(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Users = users
	resp.PageParam = pageParam
	return resp, nil
}

// UpdateUser implements the ServiceImpl interface.
func (s *ServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {
	resp = user.NewUpdateUserResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr.WithError(err))
		return resp, nil
	}

	err = service.NewUpdateUserService(ctx).UpdateUser(req)
	return resp, err
}
