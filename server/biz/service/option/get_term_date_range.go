package option

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type GetTermDateRangeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetTermDateRangeService(ctx context.Context, RequestContext *app.RequestContext) *GetTermDateRangeService {
	return &GetTermDateRangeService{RequestContext: RequestContext, Context: ctx}
}

// Run req should not be nil and resp should not be nil
func (h *GetTermDateRangeService) Run(req *model.GetTermDateRangeRequest, resp *model.GetTermDateRangeResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in GetTermDateRangeService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.GetTermDateRangeResponse)
	}
	// todo edit your code
	return
}
