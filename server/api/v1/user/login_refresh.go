package user

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/api/v1/types"
	"meetplan/model/query"
	"meetplan/pkg/jwt"
)

type RefreshTokenReq struct {
	RefreshToken string `json:"refreshToken"`
}

func RefreshToken(ctx context.Context, c *app.RequestContext, req *RefreshTokenReq) (*LoginRes, *types.PageInfo, error) {
	claims, err := jwt.VerifyJwt(req.RefreshToken)
	if err != nil {
		return nil, nil, err
	}
	user, err := query.UserColl.FindByIDStr(ctx, claims.GetUserID())
	if err != nil {
		return nil, nil, err
	}
	accessToken, err := jwt.NewJwt(user, accessTokenExpireTime)
	if err != nil {
		return nil, nil, err
	}

	expirationTime, err := claims.GetExpirationTime()
	if err != nil {
		return nil, nil, err
	}
	refreshToken := req.RefreshToken
	expireTime := expirationTime.Time
	if time.Now().Add(7 * 24 * time.Hour).After(expirationTime.Time) {
		refreshToken, err = jwt.NewJwt(user, refreshTokenExpireTime)
		if err != nil {
			return nil, nil, err
		}
		expireTime = time.Now().Add(refreshTokenExpireTime)
	}
	return &LoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expires:      expireTime,
	}, nil, nil
}
