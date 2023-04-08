// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOrder = "orders"

// Order mapped from table <orders>
type Order struct {
	ID        int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Status    int8    `gorm:"column:status;not null" json:"status"`
	Message   *string `gorm:"column:message" json:"message"`
	PlanID    int64   `gorm:"column:plan_id;not null" json:"plan_id"`
	StudentID int64   `gorm:"column:student_id;not null" json:"student_id"`
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}
