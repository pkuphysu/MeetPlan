package service

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestListFriendLink(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/friendlink", ListFriendLink)
	w := ut.PerformRequest(h.Engine, "GET", "/api/v1/friendlink", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestCreateFriendLink(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/friendlink", CreateFriendLink)
	w := ut.PerformRequest(h.Engine, "POST", "/api/v1/friendlink", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestListUpdateRecord(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/updaterecord", ListUpdateRecord)
	w := ut.PerformRequest(h.Engine, "GET", "/api/v1/updaterecord", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestCreateUpdateRecord(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/updaterecord", CreateUpdateRecord)
	w := ut.PerformRequest(h.Engine, "POST", "/api/v1/updaterecord", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetTermDateRange(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/termdate", GetTermDateRange)
	w := ut.PerformRequest(h.Engine, "GET", "/api/v1/termdate", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUpdateTermDateRange(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/termdate", UpdateTermDateRange)
	w := ut.PerformRequest(h.Engine, "PUT", "/api/v1/termdate", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
