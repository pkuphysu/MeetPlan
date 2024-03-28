package user

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

type ListRequest struct {
	Search       string `query:"search"`
	IsActive     *bool  `query:"is_active"`
	IsTeacher    *bool  `query:"is_teacher"`
	DepartmentID string `query:"department_id"`
	MajorID      string `query:"major_id"`
	GradeID      string `query:"grade_id"`

	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}

func GetUserList(ctx context.Context, c *app.RequestContext, req *ListRequest) ([]*model.User, *types.PageInfo, error) {
	filter := bson.M{}

	if req.IsActive != nil {
		filter["isActive"] = *req.IsActive
	}
	if req.IsTeacher != nil {
		filter["isTeacher"] = *req.IsTeacher
	}
	if req.Search != "" {
		filter["$or"] = []bson.M{
			{"username": bson.M{"$regex": req.Search}},
			{"pkuID": bson.M{"$regex": req.Search}},
		}
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

	return users, pageInfo, nil
}
