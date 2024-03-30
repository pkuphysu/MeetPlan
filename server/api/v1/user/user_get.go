package user

import (
	"context"

	"meetplan/pkg/jwt"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type GetUserRequest struct {
	ID string `path:"id"`
}

func GetUser(ctx context.Context, c *app.RequestContext, req *GetUserRequest) (*model.User, *types.PageInfo, error) {
	user, err := query.UserColl.FindByIDStr(ctx, req.ID)
	if err != nil {
		return nil, nil, err
	}
	FulfillUser(ctx, user)
	return user, nil, nil
}

func GetSelf(ctx context.Context, c *app.RequestContext, _ *any) (*model.User, *types.PageInfo, error) {
	user := ctx.Value(jwt.CtxUserKey).(*model.User)
	FulfillUser(ctx, user)
	return user, nil, nil
}

func FulfillUser(ctx context.Context, user *model.User) {
	if !user.DepartmentID.IsZero() {
		department, err := departmentCache.Get(ctx, user.DepartmentID.Hex())
		if err == nil {
			user.Department = department.Department
		}
	}
	if !user.MajorID.IsZero() {
		major, err := majorCache.Get(ctx, user.MajorID.Hex())
		if err == nil {
			user.Major = major.Major
		}
	}
	if !user.GradeID.IsZero() {
		grade, err := gradeCache.Get(ctx, user.GradeID.Hex())
		if err == nil {
			user.Grade = grade.Grade
			user.IsGraduated = grade.IsGraduated
		}
	}
}
