package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type ListFriendLinkService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListFriendLinkService(ctx context.Context, RequestContext *app.RequestContext) *ListFriendLinkService {
	return &ListFriendLinkService{RequestContext: RequestContext, Context: ctx}
}

func (h *ListFriendLinkService) Run(req *model.ListFriendLinkRequest, resp *model.ListFriendLinkResponse) (err *errno.Err) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
