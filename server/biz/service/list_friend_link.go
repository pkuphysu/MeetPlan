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

// Run req should not be nil and resp should not be nil
func (h *ListFriendLinkService) Run(req *model.ListFriendLinkRequest, resp *model.ListFriendLinkResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.ListFriendLinkResponse)
	}
	// todo edit your code
	return
}
