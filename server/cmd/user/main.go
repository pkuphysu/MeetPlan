package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
	"meetplan/config"
	user "meetplan/kitex_gen/pkuphy/meetplan/user/service"
	"meetplan/pkg/constants"
)

func main() {
	config.InitDB()

	reg, err := config.NewRegistry()
	if err != nil {
		log.Fatal(err.Error())
	}
	svr := user.NewServer(new(ServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.UserRpcServiceName,
			Method:      "",
			Tags:        nil,
		}),
		server.WithRegistry(reg),
	)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
