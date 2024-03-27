package user

import (
	"context"
	"errors"
	"strings"

	"github.com/samber/lo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/cloudwego/hertz/pkg/app"

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

	PageSize int    `query:"page_size"`
	Before   string `query:"before"`
	After    string `query:"after"`
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
	if req.After != "" {
		oid, err := primitive.ObjectIDFromHex(req.After)
		if err != nil {
			return nil, nil, err
		}
		filter["_id"] = bson.M{"$lt": oid}
	} else if req.Before != "" {
		oid, err := primitive.ObjectIDFromHex(req.Before)
		if err != nil {
			return nil, nil, err
		}
		filter["_id"] = bson.M{"$gt": oid}
	}
	if req.PageSize == 0 {
		req.PageSize = -1
	}

	asc := req.After == "" && req.Before != ""
	users, err := query.UserColl.FindOffset(ctx, filter, req.PageSize, asc)
	if err != nil {
		return nil, nil, err
	}
	if len(users) == 0 {
		return nil, nil, errors.New("no more users")
	}
	if asc {
		users = lo.Reverse(users)
	}

	pageInfo := &types.PageInfo{
		Total:   0,
		HasPrev: false,
		HasNext: false,
	}

	delete(filter, "_id")
	if req.PageSize == -1 && req.After == "" && req.Before == "" {
		pageInfo.Total = len(users)
	} else {
		pageInfo.Total, err = query.UserColl.Count(ctx, filter)
		if err != nil {
			return nil, nil, err
		}
	}

	filter["_id"] = bson.M{"$gte": users[0].ID}
	exist, err := query.UserColl.Exists(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	pageInfo.HasPrev = exist
	filter["_id"] = bson.M{"$lte": users[len(users)-1].ID}
	exist, err = query.UserColl.Exists(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	pageInfo.HasNext = exist

	return users, pageInfo, nil
}
