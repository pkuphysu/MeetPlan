package service

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestLogin(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/login", Login)
	w := ut.PerformRequest(h.Engine, "GET", "/api/v1/login", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetSelf(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/user/self", GetSelf)
	w := ut.PerformRequest(h.Engine, "GET", "/api/v1/user/self", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetUser(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/user/:id", GetUser)
	w := ut.PerformRequest(h.Engine, "GET", "/api/v1/user/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestListUser(t *testing.T) {
	h := server.Default()
	h.GET("/api/v1/user", ListUser)
	w := ut.PerformRequest(h.Engine, "GET", "/api/v1/user", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
