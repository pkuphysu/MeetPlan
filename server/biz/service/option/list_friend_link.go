package option

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type ListFriendLinkService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IFriendLinkDo
}

func NewListFriendLinkService(ctx context.Context, RequestContext *app.RequestContext) *ListFriendLinkService {
	return &ListFriendLinkService{RequestContext: RequestContext, Context: ctx, DAO: query.FriendLink.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *ListFriendLinkService) Run(req *model.ListFriendLinkRequest, resp *model.ListFriendLinkResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in ListFriendLinkService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.ListFriendLinkResponse)
	}
	friendLinks, e := h.DAO.Find()
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = make([]*model.FriendLink, len(friendLinks))
	for i, friendLink := range friendLinks {
		desc := ""
		if friendLink.Description != nil {
			desc = *friendLink.Description
		}
		resp.Data[i] = &model.FriendLink{
			Name:        friendLink.Name,
			Url:         friendLink.URL,
			Description: desc,
		}
	}
	return
}
