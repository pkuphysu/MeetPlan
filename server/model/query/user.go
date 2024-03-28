package query

import "meetplan/model"

var (
	UserColl       = New[model.User]("user")
	GradeColl      = New[model.Grade]("grade")
	MajorColl      = New[model.Major]("major")
	DepartmentColl = New[model.Department]("department")
)
