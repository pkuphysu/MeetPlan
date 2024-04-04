package main

import (
	"meetplan/api"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()
	h.Static("/file", "./")

	api.RegisterRoutes(h.Group("/api"))

	h.Spin()
}
