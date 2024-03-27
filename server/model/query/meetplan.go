package query

import (
	"context"
	"errors"
	"maps"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"meetplan/model"
)

var (
	ResourceVersionOutdated = errors.New("resource version outdated")
	MeetPlanColl            = New[model.MeetPlan]("meetplan")
)

func UpdateMeetPlan(ctx context.Context, meetplan *model.MeetPlan, update bson.M) (*model.MeetPlan, error) {
	set := bson.M{
		"updatedAt": time.Now(),
	}
	maps.Copy(set, update)

	res, err := MeetPlanColl.Raw().UpdateOne(ctx,
		bson.M{"_id": meetplan.ID, "resourceVersion": meetplan.ResourceVersion},
		bson.M{"$inc": bson.M{"resourceVersion": 1}, "$set": set})
	if err != nil {
		return nil, err
	}
	if res.MatchedCount != 1 {
		return nil, ResourceVersionOutdated
	}
	return MeetPlanColl.FindByID(ctx, meetplan.ID)
}

func AddMeetPlanOrder(ctx context.Context, meetplan *model.MeetPlan, order *model.MeetPlanOrder) (*model.MeetPlan, error) {
	now := time.Now()
	order.ID = primitive.NewObjectID()
	order.CreatedAt = now
	order.UpdatedAt = now

	res, err := MeetPlanColl.Raw().UpdateOne(ctx,
		bson.M{"_id": meetplan.ID, "resourceVersion": meetplan.ResourceVersion},
		bson.M{
			"$inc":  bson.M{"resourceVersion": 1},
			"$set":  bson.M{"updatedAt": now},
			"$push": bson.M{"orders": order},
		},
	)
	if err != nil {
		return nil, err
	}
	if res.MatchedCount != 1 {
		return nil, ResourceVersionOutdated
	}
	return MeetPlanColl.FindByID(ctx, meetplan.ID)
}

func UpdateMeetPlanOrder(ctx context.Context, meetplan *model.MeetPlan, order *model.MeetPlanOrder, update bson.M) (*model.MeetPlan, error) {
	now := time.Now()
	set := bson.M{
		"updatedAt":          now,
		"orders.$.updatedAt": now,
	}
	for k, v := range update {
		set["orders.$."+k] = v
	}

	res, err := MeetPlanColl.Raw().UpdateOne(ctx,
		bson.M{
			"_id":             meetplan.ID,
			"orders._id":      order.ID,
			"resourceVersion": meetplan.ResourceVersion,
		},
		bson.M{
			"$inc": bson.M{"resourceVersion": 1},
			"$set": set,
		},
	)
	if err != nil {
		return nil, err
	}
	if res.MatchedCount != 1 {
		return nil, ResourceVersionOutdated
	}
	return MeetPlanColl.FindByID(ctx, meetplan.ID)
}
