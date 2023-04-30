package service

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"

	model "meetplan/biz/model"
)

func TestDeleteMeetPlanService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewDeleteMeetPlanService(ctx, c)
	// init req and assert value
	req := &model.DeleteMeetPlanRequest{}
	resp := &model.DeleteMeetPlanResponse{}
	err := s.Run(req, resp)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
