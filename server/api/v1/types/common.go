package types

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route"
)

type PageInfo struct {
	Total   int  `json:"total"`
	HasPrev bool `json:"has_prev"`
	HasNext bool `json:"has_next"`
}

type Response[T any] struct {
	Code     int       `json:"code"`
	Data     T         `json:"data,omitempty"`
	PageInfo *PageInfo `json:"page_info,omitempty"`
	Error    string    `json:"error,omitempty"`
}

func register[T, R any](h *route.RouterGroup, method string, path string,
	f func(ctx context.Context, c *app.RequestContext, req *T) (R, *PageInfo, error)) {

	handler := func(ctx context.Context, c *app.RequestContext) {
		req := new(T)
		if err := c.BindAndValidate(req); err != nil {
			_ = c.Error(err)
			c.JSON(400, Response[R]{Error: err.Error()})
			return
		}
		res, pageInfo, err := f(ctx, c, req)
		if err != nil {
			_ = c.Error(err)
			c.JSON(400, Response[R]{Error: err.Error()})
			return
		}
		c.JSON(200, Response[R]{
			Data:     res,
			PageInfo: pageInfo,
		})
	}
	path = strings.TrimRight(path, "/")
	h.Handle(method, path, handler)
	h.Handle(method, path+"/", handler)
}

func RegisterGet[T, R any](h *route.RouterGroup, path string,
	f func(ctx context.Context, c *app.RequestContext, req *T) (R, *PageInfo, error)) {
	register(h, consts.MethodGet, path, f)
}

func RegisterPost[T, R any](h *route.RouterGroup, path string,
	f func(ctx context.Context, c *app.RequestContext, req *T) (R, *PageInfo, error)) {
	register(h, consts.MethodPost, path, f)
}

func RegisterPut[T, R any](h *route.RouterGroup, path string,
	f func(ctx context.Context, c *app.RequestContext, req *T) (R, *PageInfo, error)) {
	register(h, consts.MethodPut, path, f)
}

func RegisterDelete[T, R any](h *route.RouterGroup, path string,
	f func(ctx context.Context, c *app.RequestContext, req *T) (R, *PageInfo, error)) {
	register(h, consts.MethodDelete, path, f)
}
