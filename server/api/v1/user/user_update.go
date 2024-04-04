package user

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type UpdateUserRequest struct {
	ID              string        `path:"id"`
	IsActive        *bool         `json:"isActive"`
	IsAdmin         *bool         `json:"isAdmin"`
	IsTeacher       *bool         `json:"isTeacher"`
	Name            *string       `json:"name"`
	PkuID           *string       `json:"pkuID"` // pkuID 为空说明是外校师生
	Email           *string       `json:"email"`
	EmailConfirming *string       `json:"emailConfirming"`
	PhoneNumber     *string       `json:"phoneNumber"`
	Gender          *model.Gender `json:"gender"`
	Birthday        *string       `json:"birthday"`
	Avatar          *string       `json:"avatar"`
	DepartmentID    *string       `json:"departmentID"`
	Office          *string       `json:"office"`
	Introduction    *string       `json:"introduction"`
	Dorm            *string       `json:"dorm"`
	MajorID         *string       `json:"majorID"`
	GradeID         *string       `json:"gradeID"`
}

func UpdateUser(ctx context.Context, c *app.RequestContext, req *UpdateUserRequest) (*model.User, *types.PageInfo, error) {
	dbUser, err := query.UserColl.FindByIDStr(ctx, req.ID)
	if err != nil {
		return nil, nil, err
	}
	dbUser.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	if req.IsActive != nil {
		dbUser.IsActive = *req.IsActive
	}
	if req.IsAdmin != nil {
		dbUser.IsAdmin = *req.IsAdmin
	}
	if req.IsTeacher != nil {
		dbUser.IsTeacher = *req.IsTeacher
	}
	if req.Name != nil {
		dbUser.Name = *req.Name
	}
	if req.PkuID != nil {
		dbUser.PkuID = *req.PkuID
	}
	if req.Email != nil {
		dbUser.Email = *req.Email
	}
	if req.EmailConfirming != nil {
		dbUser.EmailConfirming = *req.EmailConfirming
	}
	if req.PhoneNumber != nil {
		dbUser.PhoneNumber = *req.PhoneNumber
	}
	if req.Gender != nil {
		dbUser.Gender = *req.Gender
	}
	if req.Birthday != nil {
		dbUser.Birthday = *req.Birthday
	}
	if req.Avatar != nil {
		dbUser.Avatar = *req.Avatar
	}
	if req.DepartmentID != nil && *req.DepartmentID != primitive.NilObjectID.Hex() {
		department, err := query.DepartmentColl.FindByIDStr(ctx, *req.DepartmentID)
		if err != nil {
			return nil, nil, err
		}
		dbUser.DepartmentID = department.ID
	}
	if req.Office != nil {
		dbUser.Office = *req.Office
	}
	if req.Introduction != nil {
		dbUser.Introduction = *req.Introduction
	}
	if req.Dorm != nil {
		dbUser.Dorm = *req.Dorm
	}
	if req.MajorID != nil && *req.MajorID != primitive.NilObjectID.Hex() {
		major, err := query.MajorColl.FindByIDStr(ctx, *req.MajorID)
		if err != nil {
			return nil, nil, err
		}
		dbUser.MajorID = major.ID
	}
	if req.GradeID != nil && *req.GradeID != primitive.NilObjectID.Hex() {
		grade, err := query.GradeColl.FindByIDStr(ctx, *req.GradeID)
		if err != nil {
			return nil, nil, err
		}
		dbUser.GradeID = grade.ID
	}

	err = query.UserColl.Upsert(ctx, dbUser)
	if err != nil {
		return nil, nil, err
	}
	return dbUser, nil, nil
}
