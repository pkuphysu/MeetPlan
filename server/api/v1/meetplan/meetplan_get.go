package meetplan

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type GetMeetPlanRequest struct {
	ID string `path:"id"`
}

func GetMeetPlan(ctx context.Context, c *app.RequestContext, req *GetMeetPlanRequest) (*model.MeetPlan, *types.PageInfo, error) {
	meetplan, err := query.MeetPlanColl.FindByIDStr(ctx, req.ID)
	if err != nil {
		return nil, nil, err
	}
	userIDs := make([]primitive.ObjectID, 0, len(meetplan.Orders)+1)
	userIDs = append(userIDs, meetplan.TeacherID)
	for _, order := range meetplan.Orders {
		userIDs = append(userIDs, order.StudentID)
	}
	users, err := query.UserColl.FindAll(ctx, bson.M{"_id": bson.M{"$in": userIDs}})
	if err != nil {
		return nil, nil, err
	}
	userMap := lo.SliceToMap(users, func(item *model.User) (primitive.ObjectID, *model.User) {
		return item.ID, item
	})
	meetplan.Teacher = userMap[meetplan.TeacherID]
	for _, order := range meetplan.Orders {
		order.Student = userMap[order.StudentID]
	}
	return meetplan, nil, nil
}
