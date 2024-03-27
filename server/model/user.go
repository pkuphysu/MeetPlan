package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt       primitive.DateTime `bson:"createdAt" json:"created_at"`
	UpdatedAt       primitive.DateTime `bson:"updatedAt" json:"updated_at"`
	IsActive        bool               `bson:"isActive" json:"is_active"`
	IsAdmin         bool               `bson:"isAdmin" json:"is_admin"`
	IsTeacher       bool               `bson:"isTeacher" json:"is_teacher"`
	Name            string             `bson:"name" json:"name"`
	PkuID           string             `bson:"pkuID" json:"pku_id"` // pkuID 为空说明是外校师生
	Email           string             `bson:"email" json:"email"`
	EmailConfirming string             `bson:"emailConfirming" json:"email_confirming"`
	PhoneNumber     string             `bson:"phoneNumber" json:"phone_number"`
	Gender          Gender             `bson:"gender" json:"gender"`
	Birthday        string             `bson:"birthday" json:"birthday"`
	Avatar          string             `bson:"avatar" json:"avatar"`
	DepartmentID    primitive.ObjectID `bson:"departmentID" json:"department_id"`
	Department      string             `bson:"-" json:"department"`
	Office          string             `bson:"office" json:"office"`
	Introduction    string             `bson:"introduction" json:"introduction"`
	Dorm            string             `bson:"dorm" json:"dorm"`
	MajorID         primitive.ObjectID `bson:"majorID" json:"major_id"`
	Major           string             `bson:"-" json:"major"`
	GradeID         primitive.ObjectID `bson:"gradeID" json:"grade_id"`
	Grade           string             `bson:"-" json:"grade"`
	IsGraduated     bool               `bson:"-" json:"is_graduated"`
}

type Gender string

const (
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
	GenderUnknown Gender = ""
)

type Grade struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Grade       string             `bson:"grade" json:"grade"`
	IsGraduated bool               `bson:"isGraduated" json:"is_graduated"`
}

type Major struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Major string             `bson:"major" json:"major"`
}

type Department struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Department string             `bson:"department" json:"department"`
}
