package service

import (
	"github.com/pkuphysu/meetplan/gorm_gen/model"
	"github.com/pkuphysu/meetplan/kitex_gen/user"
	"github.com/samber/lo"
)

func packUsers(users []*model.User) []*user.User {
	var resp []*user.User
	for _, u := range users {
		resp = append(resp, packUser(u))
	}
	return resp
}

func packUser(u *model.User) *user.User {
	return &user.User{
		Id:         &u.ID,
		PkuId:      &u.PkuID,
		Name:       &u.Name,
		Email:      &u.Email,
		IsActive:   &u.IsActive,
		IsTeacher:  &u.IsTeacher,
		IsAdmin:    &u.IsAdmin,
		Gender:     user.GenderPtr(lo.If(u.Gender, user.Gender_Female).Else(user.Gender_Male)),
		Avatar:     u.Avatar,
		Department: u.Department,
		Phone:      u.Phone,
		Major:      u.Major,
		Grade:      u.Grade,
		Dorm:       u.Dorm,
		Office:     u.Office,
		Introduce:  u.Introduce,
	}
}
