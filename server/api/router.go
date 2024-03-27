package api

import (
	"github.com/cloudwego/hertz/pkg/route"

	v1 "meetplan/api/v1"
)

func RegisterRoutes(r *route.RouterGroup) {
	// Register your routes here.
	v1.RegisterRoutes(r.Group("/v1"))
}
