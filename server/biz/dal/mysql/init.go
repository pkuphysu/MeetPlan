package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"meetplan/biz/gorm_gen/query"
	"meetplan/config"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(config.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	query.SetDefault(DB)
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Printf("get sqlDB error - %v\n", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		fmt.Printf("close sqlDB error - %v\n", err)
		return
	}
}
