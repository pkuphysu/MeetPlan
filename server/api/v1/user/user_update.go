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
	IsActive        *bool         `json:"is_active"`
	IsAdmin         *bool         `json:"is_admin"`
	IsTeacher       *bool         `json:"is_teacher"`
	Name            *string       `json:"name"`
	PkuID           *string       `json:"pku_id"` // pkuID 为空说明是外校师生
	Email           *string       `json:"email"`
	EmailConfirming *string       `json:"email_confirming"`
	PhoneNumber     *string       `json:"phone_number"`
	Gender          *model.Gender `json:"gender"`
	Birthday        *string       `json:"birthday"`
	Avatar          *string       `json:"avatar"`
	DepartmentID    *string       `json:"department_id"`
	Office          *string       `json:"office"`
	Introduction    *string       `json:"introduction"`
	Dorm            *string       `json:"dorm"`
	MajorID         *string       `json:"major_id"`
	GradeID         *string       `json:"grade_id"`
}

func UpdateUser(ctx context.Context, c *app.RequestContext, req *UpdateUserRequest) (*model.User, *types.PageInfo, error) {
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, nil, err
	}
	dbUser, err := query.UserColl.FindByID(ctx, id)
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

	err = query.UserColl.Upsert(ctx, dbUser)
	if err != nil {
		return nil, nil, err
	}
	return dbUser, nil, nil
}
