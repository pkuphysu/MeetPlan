package main

import (
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var dataMap = map[string]func(columnType gorm.ColumnType) (dataType string){
	// bool mapping
	"tinyint": func(columnType gorm.ColumnType) (dataType string) {
		detailType, _ := columnType.ColumnType()
		if strings.HasPrefix(detailType, "tinyint(1)") {
			return "bool"
		}
		return "int8"
	},
	"json": func(columnType gorm.ColumnType) (dataType string) {
		return "datatypes.JSON"
	},
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./biz/gorm_gen/query",
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

	order := g.GenerateModel("orders", gen.FieldIgnore("create_time", "update_time"))
	g.ApplyBasic(
		g.GenerateModel("users", gen.FieldIgnore("create_time", "update_time")),
		g.GenerateModel("plans",
			gen.FieldIgnore("create_time", "update_time"),
			gen.FieldRelate(field.HasMany, "Orders", order, &field.RelateConfig{
				RelateSlicePointer: true,
			}),
		),
		order,
		g.GenerateModel("plan_view",
			gen.FieldIgnore("create_time", "update_time"),
			gen.FieldType("is_valid", "bool"),
			gen.FieldType("quota_left", "int8"),
		),
		g.GenerateModel("options", gen.FieldIgnore("create_time", "update_time")),
	)
	g.Execute()
}
