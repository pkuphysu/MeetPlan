package main

import (
	"context"
	"github.com/pkuphysu/meetplan/cmd/user/service"
	user "github.com/pkuphysu/meetplan/kitex_gen/user"
	"github.com/pkuphysu/meetplan/pkg/errno"
)

// AuthImpl implements the last service interface defined in the IDL.
type AuthImpl struct{}

// Login implements the AuthImpl interface.
func (s *AuthImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the AuthImpl interface.
func (s *AuthImpl) GetUser(ctx context.Context, req *user.GetUserReq) (resp *user.GetUserResp, err error) {
	resp = user.NewGetUserResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	u, err := service.NewGetUserService(ctx).GetUser(ctx, req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.User = u
	return resp, nil
}

// MGetUser implements the AuthImpl interface.
func (s *AuthImpl) MGetUser(ctx context.Context, req *user.MGetUserReq) (resp *user.MGetUserResp, err error) {
	resp = user.NewMGetUserResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewMGetUserService(ctx).MGetUser(ctx, req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
}

// UpdateUser implements the AuthImpl interface.
func (s *AuthImpl) UpdateUser(ctx context.Context, req *user.UpdateUserReq) (resp *user.UpdateUserResp, err error) {
	resp = user.NewUpdateUserResp()
	if err = req.IsValid(); err != nil {
		resp.BaseResp = errno.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewUpdateUserService(ctx).UpdateUser(ctx, req)
	return
}
