package meetplan

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type ListMeetPlanRequest struct {
	TeacherIDs string `json:"teacher_ids"`
	StudentIDs string `json:"student_ids"`
	StartTime  int    `json:"start_time"`
	EndTime    int    `json:"end_time"`
}

func ListMeetPlan(ctx context.Context, c *app.RequestContext, req *ListMeetPlanRequest) ([]*model.MeetPlan, *types.PageInfo, error) {
	filter := bson.M{}
	if req.TeacherIDs != "" {
		var teaIDs []primitive.ObjectID
		for _, id := range strings.Split(req.TeacherIDs, ",") {
			teaID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, nil, err
			}
			teaIDs = append(teaIDs, teaID)
		}
		filter["teacherID"] = bson.M{"$in": teaIDs}
	}
	if req.StudentIDs != "" {
		var stuIDs []primitive.ObjectID
		for _, id := range strings.Split(req.StudentIDs, ",") {
			stuID, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, nil, err
			}
			stuIDs = append(stuIDs, stuID)
		}
		filter["orders.studentID"] = bson.M{"$in": stuIDs}

	}
	if req.StartTime > 0 {
		filter["startTime"] = bson.M{"$gte": req.StartTime}
	}
	if req.EndTime > 0 {
		filter["endTime"] = bson.M{"$lte": req.EndTime}
	}

	meetplans, err := query.MeetPlanColl.FindOffset(ctx, filter, 0, -1)
	if err != nil {
		return nil, nil, err
	}
	return meetplans, nil, nil
}
