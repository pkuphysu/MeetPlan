package option

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/biz/gorm_gen"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type CreateFriendLinkService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IFriendLinkDo
}

func NewCreateFriendLinkService(ctx context.Context, RequestContext *app.RequestContext) *CreateFriendLinkService {
	return &CreateFriendLinkService{RequestContext: RequestContext, Context: ctx, DAO: query.FriendLink.WithContext(ctx)}
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
	e := h.DAO.Create(&gorm_gen.FriendLink{
		Name:        req.Name,
		URL:         req.Url,
		Description: &req.Description,
	})
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = &model.FriendLink{
		Name:        req.Name,
		Url:         req.Url,
		Description: req.Description,
	}
	return
}
