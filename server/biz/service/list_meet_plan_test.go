package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	model "meetplan/biz/model"
	"testing"
)

func TestListMeetPlanService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewListMeetPlanService(ctx, c)
	// init req and assert value
	req := &model.ListMeetPlanRequest{}
	resp := &model.ListMeetPlanResponse{}
	err := s.Run(req, resp)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
