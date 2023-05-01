package user

import (
	"context"
	"errors"

	"meetplan/internal/jwt"

	"github.com/samber/lo"

	"gorm.io/gorm"

	"meetplan/biz/gorm_gen"

	"meetplan/biz/gorm_gen/query"

	"meetplan/internal/oidc_rp"

	"github.com/cloudwego/hertz/pkg/app"

	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	UserDAO        query.IUserDo
}

func NewLoginService(ctx context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: ctx, UserDAO: query.User.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *LoginService) Run(req *model.LoginRequest, resp *model.LoginResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in LoginService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.LoginResponse)
	}

	token, e := oidc_rp.GetOauth2Config().Exchange(h.Context, req.Code)
	if e != nil {
		return errno.ToDependencyErr(e)
	}

	userInfo, e := oidc_rp.GetProvider().UserInfo(h.Context, oidc_rp.GetOauth2Config().TokenSource(h.Context, token))
	if e != nil {
		return errno.ToDependencyErr(e)
	}

	var claims oidc_rp.Claims
	if e := userInfo.Claims(&claims); e != nil {
		return errno.ToInternalErr(e)
	}
	if claims.PkuID == "" {
		return errno.NewInternalErr("no pku_id in claims")
	}
	if !claims.IsPku {
		return errno.NewInternalErr("not a pku user")
	}

	user, e := h.UserDAO.Where(query.User.PkuID.Eq(claims.PkuID)).First()
	if e != nil {
		if !errors.Is(e, gorm.ErrRecordNotFound) {
			return errno.ToInternalErr(e)
		}
		user = &gorm_gen.User{
			ID:           0,
			PkuID:        claims.PkuID,
			Name:         claims.Name,
			Email:        claims.Email,
			EmailChange:  nil,
			IsActive:     true,
			IsTeacher:    claims.IsTeacher,
			IsAdmin:      false,
			Gender:       nil,
			Avatar:       nil,
			Department:   lo.If(claims.Department != "", &claims.Department).Else(nil),
			Phone:        lo.If(claims.PhoneNumber != "", &claims.PhoneNumber).Else(nil),
			Major:        nil,
			Grade:        nil,
			Dorm:         nil,
			Office:       lo.If(claims.Address.Formatted != "", &claims.Address.Formatted).Else(nil),
			Introduction: lo.If(claims.Introduce != "", &claims.Introduce).Else(nil),
		}
		if e := h.UserDAO.Create(user); e != nil {
			return errno.ToInternalErr(e)
		}
	}

	jwtToken, e := jwt.NewJwt(user)
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Jwt = jwtToken
	return
}
