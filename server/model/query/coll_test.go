package query

import (
	"reflect"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"meetplan/model"
)

func TestReflect(t *testing.T) {
	u := &model.Department{}
	u.ID = primitive.NewObjectID()
	u.Department = "计算机科学与技术系"
	assert.Assert(t, reflect.ValueOf(u).Elem().FieldByName("ID").Interface().(primitive.ObjectID).Hex() == u.ID.Hex())
}
