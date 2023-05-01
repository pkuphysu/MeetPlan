package user

import (
	"context"

	"meetplan/biz/dal/pack"
	"meetplan/biz/gorm_gen"

	"github.com/cloudwego/hertz/pkg/app"

	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type GetSelfService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetSelfService(ctx context.Context, RequestContext *app.RequestContext) *GetSelfService {
	return &GetSelfService{RequestContext: RequestContext, Context: ctx}
}

// Run req should not be nil and resp should not be nil
func (h *GetSelfService) Run(req *model.GetSelfRequest, resp *model.GetSelfResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in GetSelfService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.GetSelfResponse)
	}

	if u, ok := h.RequestContext.Get("user"); ok {
		resp.Data = pack.UserDal2Vo(u.(*gorm_gen.User))
	} else {
		err = errno.NewInternalErr("get user from context failed")
	}
	return
}
