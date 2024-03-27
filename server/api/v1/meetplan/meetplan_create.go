package meetplan

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type CreateMeetPlanRequest struct {
	MeetPlans []*model.MeetPlan `json:"meetplans"`
}

func CreateMeetPlans(ctx context.Context, c *app.RequestContext, req *CreateMeetPlanRequest) ([]*model.MeetPlan, *types.PageInfo, error) {
	err := query.MeetPlanColl.UpsertMany(ctx, req.MeetPlans)
	if err != nil {
		return nil, nil, err
	}
	return req.MeetPlans, nil, nil
}
