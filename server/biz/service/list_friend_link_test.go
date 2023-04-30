package service

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"

	model "meetplan/biz/model"
)

func TestListFriendLinkService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewListFriendLinkService(ctx, c)
	// init req and assert value
	req := &model.ListFriendLinkRequest{}
	resp := &model.ListFriendLinkResponse{}
	err := s.Run(req, resp)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
