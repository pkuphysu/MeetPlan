package user

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"

	"meetplan/api/v1/types"
	"meetplan/model/query"
	"meetplan/pkg/jwt"
	"meetplan/pkg/oidc"
)

type LoginReq struct {
	Code string `json:"code"`
}

type LoginRes struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	Expires      time.Time `json:"expires"`
}

const (
	accessTokenExpireTime  = time.Minute * 30
	refreshTokenExpireTime = time.Hour * 24 * 28
)

func Login(ctx context.Context, c *app.RequestContext, req *LoginReq) (*LoginRes, *types.PageInfo, error) {
	token, err := oidc.GetOauth2Config().Exchange(ctx, req.Code)
	if err != nil {
		return nil, nil, err
	}
	userInfo, err := oidc.GetProvider().UserInfo(ctx, oidc.GetOauth2Config().TokenSource(ctx, token))
	if err != nil {
		return nil, nil, err
	}
	var cliaim oidc.Claims
	if err = userInfo.Claims(&cliaim); err != nil {
		return nil, nil, err
	}
	if cliaim.PkuID == "" {
		return nil, nil, errors.New("invalid pku id")
	}
	if !cliaim.IsPku {
		return nil, nil, errors.New("not pku user")
	}

	user, err := query.UserColl.FindOne(ctx, bson.M{"pkuID": cliaim.PkuID})
	if err != nil {
		return nil, nil, err
	}

	jwtToken, err := jwt.NewJwt(user, accessTokenExpireTime)
	if err != nil {
		return nil, nil, err
	}
	refreshToken, err := jwt.NewJwt(user, refreshTokenExpireTime)
	if err != nil {
		return nil, nil, err
	}
	return &LoginRes{
		AccessToken:  jwtToken,
		RefreshToken: refreshToken,
		Expires:      time.Now().Add(refreshTokenExpireTime),
	}, nil, nil
}
