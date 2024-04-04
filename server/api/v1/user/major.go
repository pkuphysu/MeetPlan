package user

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/sync/singleflight"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/cache"
	"meetplan/model/query"
)

var (
	majorCache = cache.NewLoader[*model.Major](func(ctx context.Context, key string) (*model.Major, error) {
		id, err := primitive.ObjectIDFromHex(key)
		if err != nil {
			return nil, err
		}
		return query.MajorColl.FindOne(ctx, bson.M{"_id": id})

	}, &singleflight.Group{}, true)
)

func init() {
	majors, err := query.MajorColl.FindAll(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	for _, major := range majors {
		_ = majorCache.Set(context.Background(), major.ID.Hex(), major)
	}
}

type ListMajorRequest struct {
	Search   string `query:"search"`
	Page     int    `query:"page"`
	PageSize int    `query:"pageSize"`
}

func GetMajorList(ctx context.Context, c *app.RequestContext, req *ListMajorRequest) ([]*model.Major, *types.PageInfo, error) {
	filter := bson.M{}
	if req.Search != "" {
		filter["major"] = bson.M{"$regex": req.Search}
	}

	majors, err := query.MajorColl.FindPage(ctx, filter, req.Page, req.PageSize)
	if err != nil {
		return nil, nil, err
	}
	count, err := query.MajorColl.Count(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	for _, major := range majors {
		_ = majorCache.Set(ctx, major.ID.Hex(), major)
	}
	return majors, &types.PageInfo{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

type CreateMajorRequest struct {
	Major string `json:"major"`
}

func CreateMajor(ctx context.Context, c *app.RequestContext, req *CreateMajorRequest) (*model.Major, *types.PageInfo, error) {
	exists, err := query.MajorColl.Exists(ctx, bson.M{"major": req.Major})
	if err != nil {
		return nil, nil, err
	}
	if exists {
		return nil, nil, errors.New("major already exists")
	}
	major := &model.Major{
		ID:    primitive.NewObjectID(),
		Major: req.Major,
	}
	err = query.MajorColl.Upsert(ctx, major)
	if err != nil {
		return nil, nil, err
	}
	_ = majorCache.Set(ctx, major.ID.Hex(), major)
	return major, nil, nil
}

type UpdateMajorRequest struct {
	ID    string `path:"id"`
	Major string `json:"major"`
}

func UpdateMajor(ctx context.Context, c *app.RequestContext, req *UpdateMajorRequest) (*model.Major, *types.PageInfo, error) {
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, nil, err
	}
	exists, err := query.MajorColl.Exists(ctx, bson.M{"major": req.Major, "_id": bson.M{"$ne": id}})
	if err != nil {
		return nil, nil, err
	}
	if exists {
		return nil, nil, errors.New("major already exists")
	}
	major := &model.Major{
		ID:    id,
		Major: req.Major,
	}
	err = query.MajorColl.Upsert(ctx, major)
	if err != nil {
		return nil, nil, err
	}
	_ = majorCache.Set(ctx, major.ID.Hex(), major)
	return major, nil, nil
}
