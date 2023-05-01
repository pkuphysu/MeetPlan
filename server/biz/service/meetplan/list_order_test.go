package meetplan

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	model "meetplan/biz/model"
)

func TestListOrderService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewListOrderService(ctx, c)
	// init req and assert value
	req := &model.ListOrderRequest{}
	resp := &model.ListOrderResponse{}
	err := s.Run(req, resp)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
