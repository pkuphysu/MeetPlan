package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MeetPlan struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt       time.Time          `bson:"createdAt" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updatedAt" json:"updated_at"`
	ResourceVersion int                `bson:"resourceVersion" json:"resource_version"`
	Place           string             `bson:"place" json:"place"`
	StartTime       time.Time          `bson:"startTime" json:"start_time"`
	EndTime         time.Time          `bson:"endTime" json:"end_time"`
	Message         string             `bson:"message" json:"message"`
	TeacherID       primitive.ObjectID `bson:"teacherID" json:"teacher_id"`
	Teacher         *User              `bson:"-" json:"teacher"`
	Capacity        int                `bson:"capacity" json:"capacity"`
	Orders          []*MeetPlanOrder   `bson:"orders" json:"orders"`
}

type MeetPlanOrder struct {
	ID        primitive.ObjectID  `bson:"_id" json:"id"`
	CreatedAt time.Time           `bson:"createdAt" json:"created_at"`
	UpdatedAt time.Time           `bson:"updatedAt" json:"updated_at"`
	StudentID primitive.ObjectID  `bson:"studentID" json:"student_id"`
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
