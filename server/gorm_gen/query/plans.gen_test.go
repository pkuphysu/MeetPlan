// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkuphysu/meetplan/gorm_gen/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Plan{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Plan{}) fail: %s", err)
	}
}

func Test_planQuery(t *testing.T) {
	plan := newPlan(db)
	plan = *plan.As(plan.TableName())
	_do := plan.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(plan.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <plans> fail:", err)
		return
	}

	_, ok := plan.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from plan success")
	}

	err = _do.Create(&model.Plan{})
	if err != nil {
		t.Error("create item in table <plans> fail:", err)
	}

	err = _do.Save(&model.Plan{})
	if err != nil {
		t.Error("create item in table <plans> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Plan{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <plans> fail:", err)
	}

	_, err = _do.Select(plan.ALL).Take()
	if err != nil {
		t.Error("Take() on table <plans> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <plans> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <plans> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <plans> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Plan{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <plans> fail:", err)
	}

	_, err = _do.Select(plan.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <plans> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <plans> fail:", err)
	}

	_, err = _do.Select(plan.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <plans> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <plans> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <plans> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <plans> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Plan{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <plans> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <plans> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <plans> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <plans> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <plans> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <plans> fail:", err)
	}
}
