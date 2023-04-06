package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/pkuphysu/meetplan/config"
	order "github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/order/service"
	"github.com/pkuphysu/meetplan/pkg/constants"
	"log"
)

func main() {
	config.InitDB()

	reg, err := config.NewRegistry()
	svr := order.NewServer(new(ServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.OrderRpcServiceName,
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
