// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gorm_gen

const TableNameFriendLink = "friend_links"

// FriendLink mapped from table <friend_links>
type FriendLink struct {
	ID          int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string  `gorm:"column:name;not null" json:"name"`
	URL         string  `gorm:"column:url;not null" json:"url"`
	Description *string `gorm:"column:description" json:"description"`
}

// TableName FriendLink's table name
func (*FriendLink) TableName() string {
	return TableNameFriendLink
}