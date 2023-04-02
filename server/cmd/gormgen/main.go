package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

var dataMap = map[string]func(detailType string) (dataType string){
	// bool mapping
	"tinyint": func(detailType string) (dataType string) {
		if strings.HasPrefix(detailType, "tinyint(1)") {
			return "bool"
		}
		return "int8"
	},
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./gorm_gen/query",
		WithUnitTest:      true,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  false,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	db, _ := gorm.Open(mysql.Open("meetplan:meetplan@(127.0.0.1:3306)/meetplan?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)

	g.WithDataTypeMap(dataMap)

	g.ApplyBasic(
		g.GenerateModel("users", gen.FieldIgnore("create_time", "update_time")),
		g.GenerateModel("plans", gen.FieldIgnore("create_time", "update_time")),
		g.GenerateModel("orders", gen.FieldIgnore("create_time", "update_time")),
	)
	g.Execute()
}
