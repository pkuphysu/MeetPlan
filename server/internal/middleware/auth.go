package middleware

import (
	"context"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"meetplan/biz/gorm_gen/query"
	"meetplan/internal/jwt"
	"meetplan/pkg/constant"
)

func Jwt() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		if _, ok := ctx.Get(constant.CtxKeyUser); ok {
			return
		}
		authHeader := ctx.Request.Header.Get(constant.HeaderAuthorization)
		if authHeader == "" {
			ctx.AbortWithStatus(consts.StatusUnauthorized)
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != constant.HeaderAuthorizationBearer {
			ctx.AbortWithStatus(consts.StatusUnauthorized)
			return
		}
		token := parts[1]
		claims, err := jwt.VerifyJwt(token)
		if err != nil {
			ctx.AbortWithStatus(consts.StatusUnauthorized)
			return
		}
		user_id, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			ctx.AbortWithStatus(consts.StatusUnauthorized)
			return
		}
		user, err := query.User.WithContext(c).Where(query.User.ID.Eq(user_id)).First()
		if err != nil {
			ctx.AbortWithStatus(consts.StatusUnauthorized)
			return
		}
		ctx.Set(constant.CtxKeyUser, user)
		ctx.Next(c)
	}
}
