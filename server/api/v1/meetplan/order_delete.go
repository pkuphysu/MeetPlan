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

type DeleteOrderReq struct {
	MeetPlanID string `path:"id"`
	OrderID    string `path:"order_id"`
}

func DeleteOrder(ctx context.Context, c *app.RequestContext, req *DeleteOrderReq) (*model.MeetPlan, *types.PageInfo, error) {
	meetplan, err := query.MeetPlanColl.FindByIDStr(ctx, req.MeetPlanID)
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

	update := bson.M{
		"status": model.MeetPlanOrderStatusCancelled,
	}

	meetplan, err = query.UpdateMeetPlanOrder(ctx, meetplan, order, update)
	if err != nil {
		return nil, nil, err
	}

	return meetplan, nil, nil
}
