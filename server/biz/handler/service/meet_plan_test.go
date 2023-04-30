package service

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetMeetPlan(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/meetplan/:id", GetMeetPlan)
	w := ut.PerformRequest(h.Engine, "GET", "/api/v1/meetplan/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestListMeetPlan(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/meetplan", ListMeetPlan)
	w := ut.PerformRequest(h.Engine, "GET", "/api/v1/meetplan", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestCreateMeetPlan(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/meetplan", CreateMeetPlan)
	w := ut.PerformRequest(h.Engine, "POST", "/api/v1/meetplan", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateMeetPlan(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/meetplan/:id", UpdateMeetPlan)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/v1/meetplan/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDeleteMeetPlan(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/meetplan/:id", DeleteMeetPlan)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/v1/meetplan/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDeleteMeetPlans(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/meetplan", DeleteMeetPlans)
	w := ut.PerformRequest(h.Engine, "DELETE", "/api/v1/meetplan", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetOrder(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/order/:id", GetOrder)
	w := ut.PerformRequest(h.Engine, "GET", "/api/v1/order/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestListOrder(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/order", ListOrder)
	w := ut.PerformRequest(h.Engine, "GET", "/api/v1/order", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestCreateOrder(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/order", CreateOrder)
	w := ut.PerformRequest(h.Engine, "POST", "/api/v1/order", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateOrder(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/order/:id", UpdateOrder)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/v1/order/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestCreateMeetPlanAndOrder(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/meetplanorder", CreateMeetPlanAndOrder)
	w := ut.PerformRequest(h.Engine, "POST", "/api/v1/meetplanorder", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
