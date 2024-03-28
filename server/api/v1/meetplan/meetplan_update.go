package meetplan

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type UpdateMeetPlanRequest struct {
	ID        string     `path:"id"`
	Place     *string    `json:"place"`
	TeacherID *string    `json:"teacher_id"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	Message   *string    `json:"message"`
}

func UpdateMeetPlan(ctx context.Context, c *app.RequestContext, req *UpdateMeetPlanRequest) (*model.MeetPlan, *types.PageInfo, error) {
	meetplan, err := query.MeetPlanColl.FindByIDStr(ctx, req.ID)
	if err != nil {
		return nil, nil, err
	}
	update := bson.M{}
	if req.Place != nil {
		meetplan.Place = *req.Place
		update["place"] = *req.Place
	}
	if req.TeacherID != nil {
		teacherID, err := primitive.ObjectIDFromHex(*req.TeacherID)
		if err != nil {
			return nil, nil, err
		}
		meetplan.TeacherID = teacherID
		update["teacherID"] = teacherID
	}
	if req.StartTime != nil {
		meetplan.StartTime = *req.StartTime
		update["startTime"] = *req.StartTime
	}
	if req.EndTime != nil {
		meetplan.EndTime = *req.EndTime
		update["endTime"] = *req.EndTime
	}
	if req.Message != nil {
		meetplan.Message = *req.Message
		update["message"] = *req.Message
	}
	plan, err := query.UpdateMeetPlan(ctx, meetplan, update)
	if err != nil {
		return nil, nil, err
	}
	return plan, nil, nil
}
