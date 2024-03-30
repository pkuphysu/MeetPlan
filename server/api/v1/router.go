package v1

import (
	"github.com/cloudwego/hertz/pkg/route"

	"meetplan/api/middleware"
	"meetplan/api/v1/meetplan"
	"meetplan/api/v1/option"
	"meetplan/api/v1/types"
	"meetplan/api/v1/user"
)

// RegisterRoutes is a function that registers routes.
func RegisterRoutes(h *route.RouterGroup) {
	types.RegisterPost(h, "/login", user.Login)
	types.RegisterPost(h, "/login/refresh", user.RefreshToken)
	users := h.Group("/users", middleware.Jwt())
	{
		types.RegisterGet(users, "/", user.GetUserList)
		types.RegisterGet(users, "/self", user.GetSelf)
		types.RegisterGet(users, "/:id", user.GetUser)
		types.RegisterPost(users, "/", user.CreateUsers)
		types.RegisterPut(users, "/:id", user.UpdateUser)
	}

	departments := h.Group("/departments", middleware.Jwt())
	{
		types.RegisterGet(departments, "/", user.GetDepartmentList)
		types.RegisterPost(departments, "/", user.CreateDepartment)
		types.RegisterPut(departments, "/:id", user.UpdateDepartment)
	}

	majors := h.Group("/majors", middleware.Jwt())
	{
		types.RegisterGet(majors, "/", user.GetMajorList)
		types.RegisterPost(majors, "/", user.CreateMajor)
		types.RegisterPut(majors, "/:id", user.UpdateMajor)
	}

	grades := h.Group("/grades", middleware.Jwt())
	{
		types.RegisterGet(grades, "/", user.GetGradeList)
		types.RegisterPost(grades, "/", user.CreateGrade)
		types.RegisterPut(grades, "/:id", user.UpdateGrade)
	}

	meetplans := h.Group("/meetplans", middleware.Jwt())
	{
		types.RegisterGet(meetplans, "/", meetplan.ListMeetPlan)
		types.RegisterGet(meetplans, "/:id", meetplan.GetMeetPlan)
		types.RegisterPost(meetplans, "/", meetplan.CreateMeetPlans)
		types.RegisterPut(meetplans, "/:id", meetplan.UpdateMeetPlan)
		types.RegisterPost(meetplans, "/:id/orders", meetplan.CreateOrder)
		types.RegisterPut(meetplans, "/:id/orders/:order_id", meetplan.UpdateOrder)
		types.RegisterDelete(meetplans, "/:id/orders/:order_id", meetplan.DeleteOrder)
	}

	options := h.Group("/options", middleware.Jwt())
	{
		types.RegisterGet(options, "/", option.ListOption)
		types.RegisterPost(options, "/", option.CreateOptions)
		types.RegisterPut(options, "/:id", option.UpdateOption)
		types.RegisterDelete(options, "/:id", option.DeleteOption)
	}
}
