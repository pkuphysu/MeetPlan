package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/pkuphysu/meetplan/config"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/base"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/user"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/user/service"
	"github.com/pkuphysu/meetplan/pkg/constants"
	"github.com/pkuphysu/meetplan/pkg/errno"
	"github.com/samber/lo"
	"log"
)

var userClient service.Client

func initUserRpc() {
	rsv, err := config.NewResolver()
	if err != nil {
		log.Fatal(err)
	}
	userClient, err = service.NewClient(constants.UserRpcServiceName, client.WithResolver(rsv))
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserDetail(ctx context.Context, userID int64) (*user.User, error) {
	req := &user.GetUserReq{Id: lo.ToPtr(userID)}
	resp, err := userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp != nil && resp.BaseResp.StatusCode != base.StatusCode_SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.Message)
	}
	return resp.User, nil
}

func MGetUserDetailMap(ctx context.Context, userIDs []int64) (map[int64]*user.User, error) {
	req := &user.MGetUserReq{Ids: userIDs}
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp != nil && resp.BaseResp.StatusCode != base.StatusCode_SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.Message)
	}
	userMap := make(map[int64]*user.User)
	for _, u := range resp.Users {
		userMap[*u.Id] = u
	}
	return userMap, nil
}
