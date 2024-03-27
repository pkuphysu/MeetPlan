package meetplan

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type UpdateOrderReq struct {
	MeetPlanID string  `path:"id"`
	OrderID    string  `path:"order_id"`
	StudentID  *string `json:"student_id"`
	Message    *string `json:"message"`
	Status     *string `json:"status"`
}

func UpdateOrder(ctx context.Context, c *app.RequestContext, req *UpdateOrderReq) (*model.MeetPlan, *types.PageInfo, error) {
	id, err := primitive.ObjectIDFromHex(req.MeetPlanID)
	if err != nil {
		return nil, nil, err
	}
	meetplan, err := query.MeetPlanColl.FindByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	orderID, err := primitive.ObjectIDFromHex(req.OrderID)
	if err != nil {
		return nil, nil, err
	}
	order, ok := lo.Find(meetplan.Orders, func(order *model.MeetPlanOrder) bool {
		return order.ID == orderID
	})
	if !ok {
		return nil, nil, errors.New("order not found")
	}

	update := bson.M{}
	if req.StudentID != nil {
		studentID, err := primitive.ObjectIDFromHex(*req.StudentID)
		if err != nil {
			return nil, nil, err
		}
		order.StudentID = studentID
		update["studentID"] = studentID
	}
	if req.Message != nil {
		order.Message = *req.Message
		update["message"] = *req.Message
	}
	if req.Status != nil {
		order.Status = model.MeetPlanOrderStatus(*req.Status)
		update["status"] = *req.Status
	}

	meetplan, err = query.UpdateMeetPlanOrder(ctx, meetplan, order, update)
	if err != nil {
		return nil, nil, err
	}
	return meetplan, nil, nil
}
