package user

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	lop "github.com/samber/lo/parallel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type ListRequest struct {
	Name         string `query:"name"`
	PkuID        string `query:"pkuID"`
	IsActive     *bool  `query:"isActive"`
	IsTeacher    *bool  `query:"isTeacher"`
	IsAdmin      *bool  `query:"isAdmin"`
	DepartmentID string `query:"departmentID"`
	MajorID      string `query:"majorID"`
	GradeID      string `query:"gradeID"`

	Page     int `query:"page"`
	PageSize int `query:"pageSize"`
}

func GetUserList(ctx context.Context, c *app.RequestContext, req *ListRequest) ([]*model.User, *types.PageInfo, error) {
	filter := bson.M{}

	if req.IsActive != nil {
		filter["isActive"] = *req.IsActive
	}
	if req.IsTeacher != nil {
		filter["isTeacher"] = *req.IsTeacher
	}
	if req.IsAdmin != nil {
		filter["isAdmin"] = *req.IsAdmin
	}
	if req.Name != "" {
		filter["name"] = bson.M{"$regex": req.Name}
	}
	if req.PkuID != "" {
		filter["pkuID"] = bson.M{"$regex": req.PkuID}
	}
	if len(req.DepartmentID) > 0 {
		var ids []primitive.ObjectID
		for _, id := range strings.Split(req.DepartmentID, ",") {
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, nil, err
			}
			ids = append(ids, oid)
		}
		filter["departmentID"] = bson.M{"$in": ids}
	}
	if len(req.MajorID) > 0 {
		var ids []primitive.ObjectID
		for _, id := range strings.Split(req.MajorID, ",") {
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, nil, err
			}
			ids = append(ids, oid)
		}
		filter["majorID"] = bson.M{"$in": ids}
	}
	if len(req.GradeID) > 0 {
		var ids []primitive.ObjectID
		for _, id := range strings.Split(req.GradeID, ",") {
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return nil, nil, err
			}
			ids = append(ids, oid)
		}
		filter["gradeID"] = bson.M{"$in": ids}
	}

	users, err := query.UserColl.FindPage(ctx, filter, req.Page, req.PageSize)
	if err != nil {
		return nil, nil, err
	}
	total, err := query.UserColl.Count(ctx, filter)
	if err != nil {
		return nil, nil, err
	}

	pageInfo := &types.PageInfo{
		Page:     req.Page,
		PageSize: req.PageSize,
		Total:    total,
	}

	lop.ForEach(users, func(user *model.User, _ int) {
		FulfillUser(ctx, user)
	})

	return users, pageInfo, nil
}
