package middleware

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"meetplan/model/query"
	"meetplan/pkg/jwt"
)

func Jwt() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authHeader := c.Request.Header.Get(consts.HeaderAuthorization)
		if authHeader == "" {
			c.AbortWithStatus(consts.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.VerifyJwt(token)
		if err != nil {
			c.AbortWithStatus(consts.StatusUnauthorized)
			return
		}

		user, err := query.UserColl.FindByIDStr(ctx, claims.GetUserID())
		if err != nil {
			c.AbortWithStatus(consts.StatusUnauthorized)
			return
		}
		ctx = context.WithValue(ctx, jwt.CtxUserKey, user)
		c.Next(ctx)
	}
}
