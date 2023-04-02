package main

import (
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	user "github.com/pkuphysu/meetplan/kitex_gen/user/userservice"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(mysql.Open("meetplan:meetplan@(127.0.0.1:3306)/meetplan?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)

	svr := user.NewServer(new(AuthImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
