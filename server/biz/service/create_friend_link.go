package service

import (
	"context"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

type CreateFriendLinkService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateFriendLinkService(ctx context.Context, RequestContext *app.RequestContext) *CreateFriendLinkService {
	return &CreateFriendLinkService{RequestContext: RequestContext, Context: ctx}
}

// Run req should not be nil and resp should not be nil
func (h *CreateFriendLinkService) Run(req *model.CreateFriendLinkRequest, resp *model.CreateFriendLinkResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.CreateFriendLinkResponse)
	}
	// todo edit your code
	return
}
