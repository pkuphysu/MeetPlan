package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/pkuphysu/meetplan/config"
	plan "github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/plan/service"
	"github.com/pkuphysu/meetplan/pkg/constants"
	"log"
)

func main() {
	config.InitDB()

	reg, err := config.NewRegistry()
	if err != nil {
		log.Fatal(err.Error())
	}
	svr := plan.NewServer(new(ServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.PlanRpcServiceName,
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
