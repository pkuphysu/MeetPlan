package config

import (
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	db, err := gorm.Open(mysql.Open("meetplan:meetplan@(127.0.0.1:3306)/meetplan?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)
}
