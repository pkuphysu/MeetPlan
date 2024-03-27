package meetplan

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type CreateOrderReq struct {
	MeetPlanID string `path:"id"`
	StudentID  string `json:"student_id"`
	Message    string `json:"message"`
	Status     string `json:"status"`
}

func CreateOrder(ctx context.Context, c *app.RequestContext, req *CreateOrderReq) (*model.MeetPlan, *types.PageInfo, error) {
	id, err := primitive.ObjectIDFromHex(req.MeetPlanID)
	if err != nil {
		return nil, nil, err
	}
	meetplan, err := query.MeetPlanColl.FindByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	studentID, err := primitive.ObjectIDFromHex(req.StudentID)
	if err != nil {
		return nil, nil, err
	}

	if len(meetplan.Orders) >= meetplan.Capacity {
		return nil, nil, errors.New("meetplan is full")
	}

	order := &model.MeetPlanOrder{
		StudentID: studentID,
		Message:   req.Message,
		Status:    model.MeetPlanOrderStatus(req.Status),
	}
	meetplan, err = query.AddMeetPlanOrder(ctx, meetplan, order)
	if err != nil {
		return nil, nil, err
	}
	return meetplan, nil, nil
}
