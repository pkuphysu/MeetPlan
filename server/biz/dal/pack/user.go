package pack

import (
	"github.com/samber/lo"

	"meetplan/biz/gorm_gen"
	"meetplan/biz/model"
)

func UserDal2Vo(user *gorm_gen.User) *model.User {
	if user == nil {
		return nil
	}
	return &model.User{
		Id:        user.ID,
		PkuId:     user.PkuID,
		Name:      user.Name,
		Email:     user.Email,
		IsTeacher: user.IsTeacher,
		IsAdmin:   user.IsAdmin,
		IsActive:  user.IsActive,
		Gender: lo.IfF(user.Gender != nil, func() *model.Gender {
			if *user.Gender {
				return lo.ToPtr(model.Gender_GENDER_FEMALE)
			}
			return lo.ToPtr(model.Gender_GENDER_MALE)
		}).Else(nil),
		Avatar: lo.IfF(user.Avatar != nil, func() string {
			return *user.Avatar
		}).Else(""),
		Phone: lo.IfF(user.Phone != nil, func() string {
			return *user.Phone
		}).Else(""),
		Department: lo.IfF(user.Department != nil, func() string {
			return *user.Department
		}).Else(""),
		Major: lo.IfF(user.Major != nil, func() string {
			return *user.Major
		}).Else(""),
		Grade: lo.IfF(user.Grade != nil, func() string {
			return *user.Grade
		}).Else(""),
		Dorm: lo.IfF(user.Dorm != nil, func() string {
			return *user.Dorm
		}).Else(""),
		Office: lo.IfF(user.Office != nil, func() string {
			return *user.Office
		}).Else(""),
		Introduction: lo.IfF(user.Introduction != nil, func() string {
			return *user.Introduction
		}).Else(""),
		EmailChange: lo.IfF(user.EmailChange != nil, func() string {
			return *user.EmailChange
		}).Else(""),
	}
}
