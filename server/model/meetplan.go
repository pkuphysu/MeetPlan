package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeetPlan struct {
	ID              primitive.ObjectID `bson:"_id" json:"ID"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt" json:"updatedAt"`
	ResourceVersion int                `bson:"resourceVersion" json:"resourceVersion"`
	Place           string             `bson:"place" json:"place"`
	StartTime       time.Time          `bson:"startTime" json:"startTime"`
	EndTime         time.Time          `bson:"endTime" json:"endTime"`
	Message         string             `bson:"message" json:"message"`
	TeacherID       primitive.ObjectID `bson:"teacherID" json:"teacherID"`
	Teacher         *User              `bson:"-" json:"teacher"`
	Capacity        int                `bson:"capacity" json:"capacity"`
	Orders          []*MeetPlanOrder   `bson:"orders" json:"orders"`
}

type MeetPlanOrder struct {
	ID        primitive.ObjectID  `bson:"_id" json:"ID"`
	CreatedAt time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time           `bson:"updatedAt" json:"updatedAt"`
	StudentID primitive.ObjectID  `bson:"studentID" json:"studentID"`
	Student   *User               `bson:"-" json:"student"`
	Message   string              `bson:"message" json:"message"`
	Status    MeetPlanOrderStatus `bson:"status" json:"status"`
}

type MeetPlanOrderStatus string

const (
	MeetPlanOrderStatusCreated   MeetPlanOrderStatus = "created"
	MeetPlanOrderStatusConfirmed MeetPlanOrderStatus = "confirmed"
	MeetPlanOrderStatusCancelled MeetPlanOrderStatus = "cancelled"
	MeetPlanOrderStatusCompleted MeetPlanOrderStatus = "completed"
)
