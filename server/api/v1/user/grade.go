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
	gradeCache = cache.NewLoader[*model.Grade](func(ctx context.Context, key string) (*model.Grade, error) {
		id, err := primitive.ObjectIDFromHex(key)
		if err != nil {
			return nil, err
		}
		return query.GradeColl.FindOne(ctx, bson.M{"_id": id})
	}, &singleflight.Group{}, true)
)

func init() {
	grades, err := query.GradeColl.FindAll(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	for _, grade := range grades {
		_ = gradeCache.Set(context.Background(), grade.ID.Hex(), grade)
	}
}

type ListGradeRequest struct {
	Search string `query:"search"`
}

func GetGradeList(ctx context.Context, c *app.RequestContext, req *ListGradeRequest) ([]*model.Grade, *types.PageInfo, error) {
	filter := bson.M{}
	if req.Search != "" {
		filter["grade"] = bson.M{"$regex": req.Search}
	}

	grades, err := query.GradeColl.FindAll(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	for _, grade := range grades {
		_ = gradeCache.Set(ctx, grade.ID.Hex(), grade)
	}

	return grades, nil, nil
}

type CreateGradeRequest struct {
	Grade       string `json:"grade"`
	IsGraduated bool   `json:"is_graduated"`
}

func CreateGrade(ctx context.Context, c *app.RequestContext, req *CreateGradeRequest) (*model.Grade, *types.PageInfo, error) {
	exists, err := query.GradeColl.Exists(ctx, bson.M{"grade": req.Grade})
	if err != nil {
		return nil, nil, err
	}
	if exists {
		return nil, nil, errors.New("grade already exists")
	}
	grade := &model.Grade{
		ID:          primitive.NewObjectID(),
		Grade:       req.Grade,
		IsGraduated: req.IsGraduated,
	}
	err = query.GradeColl.Upsert(ctx, grade)
	if err != nil {
		return nil, nil, err
	}
	_ = gradeCache.Set(ctx, grade.ID.Hex(), grade)
	return grade, nil, nil
}

type UpdateGradeRequest struct {
	ID          string `path:"id"`
	Grade       string `json:"grade"`
	IsGraduated bool   `json:"is_graduated"`
}

func UpdateGrade(ctx context.Context, c *app.RequestContext, req *UpdateGradeRequest) (*model.Grade, *types.PageInfo, error) {
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, nil, err
	}
	exists, err := query.GradeColl.Exists(ctx, bson.M{"grade": req.Grade, "_id": bson.M{"$ne": id}})
	if err != nil {
		return nil, nil, err
	}
	if exists {
		return nil, nil, errors.New("grade already exists")
	}
	grade := &model.Grade{
		ID:          id,
		Grade:       req.Grade,
		IsGraduated: req.IsGraduated,
	}
	err = query.GradeColl.Upsert(ctx, grade)
	if err != nil {
		return nil, nil, err
	}
	_ = gradeCache.Set(ctx, grade.ID.Hex(), grade)
	return grade, nil, nil
}
