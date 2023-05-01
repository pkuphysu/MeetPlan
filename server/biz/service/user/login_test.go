package user

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	model "meetplan/biz/model"
)

func TestLoginService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewLoginService(ctx, c)
	// init req and assert value
	req := &model.LoginRequest{}
	resp := &model.LoginResponse{}
	err := s.Run(req, resp)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
