package service

import (
	"context"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/user"
)

type LoginServiceI interface {
	Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error)
}

func NewLoginService(ctx context.Context) LoginServiceI {
	return &loginService{}
}

type loginService struct{}

func (s *loginService) Login(ctx context.Context, req *user.LoginReq) (*user.LoginResp, error) {
	// TODO: Your code here...
	return nil, nil
}
