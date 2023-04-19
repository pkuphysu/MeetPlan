package httputil

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/cmd/api/utils/bytesconv"
)

func SendResponse(ctx context.Context, c *app.RequestContext, code int, resp interface{}) {
	if bytesconv.B2s(c.ContentType()) == "application/x-protobuf" {
		c.ProtoBuf(code, resp)
	} else {
		c.JSON(code, resp)
	}
}
