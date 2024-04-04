package user

import (
	"context"
	"errors"

	"golang.org/x/sync/singleflight"

	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/cache"
	"meetplan/model/query"
)

var (
	departmentCache = cache.NewLoader[*model.Department](func(ctx context.Context, key string) (*model.Department, error) {
		id, err := primitive.ObjectIDFromHex(key)
		if err != nil {
			return nil, err
		}
		return query.DepartmentColl.FindOne(ctx, bson.M{"_id": id})
	}, &singleflight.Group{}, true)
)

func init() {
	departments, err := query.DepartmentColl.FindAll(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	for _, department := range departments {
		_ = departmentCache.Set(context.Background(), department.ID.Hex(), department)
	}
}

type ListDepartmentRequest struct {
	Search   string `query:"search"`
	Page     int    `query:"page"`
	PageSize int    `query:"pageSize"`
}

func GetDepartmentList(ctx context.Context, c *app.RequestContext, req *ListDepartmentRequest) ([]*model.Department, *types.PageInfo, error) {
	filter := bson.M{}
	if req.Search != "" {
		filter["department"] = bson.M{"$regex": req.Search}
	}

	departments, err := query.DepartmentColl.FindPage(ctx, filter, req.Page, req.PageSize)
	if err != nil {
		return nil, nil, err
	}
	count, err := query.DepartmentColl.Count(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	for _, department := range departments {
		_ = departmentCache.Set(ctx, department.ID.Hex(), department)
	}

	return departments, &types.PageInfo{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    count,
	}, nil
}

type CreateDepartmentRequest struct {
	Department string `json:"department"`
}

func CreateDepartment(ctx context.Context, c *app.RequestContext, req *CreateDepartmentRequest) (*model.Department, *types.PageInfo, error) {
	exists, err := query.DepartmentColl.Exists(ctx, bson.M{"department": req.Department})
	if err != nil {
		return nil, nil, err
	}
	if exists {
		return nil, nil, errors.New("department already exists")
	}
	department := &model.Department{
		ID:         primitive.NewObjectID(),
		Department: req.Department,
	}
	err = query.DepartmentColl.Upsert(ctx, department)
	if err != nil {
		return nil, nil, err
	}
	_ = departmentCache.Set(ctx, department.ID.Hex(), department)
	return department, nil, nil
}

type UpdateDepartmentRequest struct {
	ID         string `path:"id"`
	Department string `json:"department"`
}

func UpdateDepartment(ctx context.Context, c *app.RequestContext, req *UpdateDepartmentRequest) (*model.Department, *types.PageInfo, error) {
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, nil, err
	}
	exists, err := query.DepartmentColl.Exists(ctx, bson.M{"department": req.Department, "_id": bson.M{"$ne": id}})
	if err != nil {
		return nil, nil, err
	}
	if exists {
		return nil, nil, errors.New("department already exists")
	}
	department := &model.Department{
		ID:         id,
		Department: req.Department,
	}
	err = query.DepartmentColl.Upsert(ctx, department)
	if err != nil {
		return nil, nil, err
	}
	_ = departmentCache.Set(ctx, department.ID.Hex(), department)
	return department, nil, nil
}
