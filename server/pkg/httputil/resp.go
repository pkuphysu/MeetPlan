package httputil

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/pkg/bytesconv"
)

func SendResponse(ctx context.Context, c *app.RequestContext, code int, resp any) {
	if bytesconv.B2s(c.ContentType()) == "application/x-protobuf" {
		c.ProtoBuf(code, resp)
	} else {
		c.JSON(code, resp)
	}
}