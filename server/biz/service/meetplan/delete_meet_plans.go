package meetplan

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/samber/lo"

	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type DeleteMeetPlansService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IPlanDo
}

func NewDeleteMeetPlansService(ctx context.Context, RequestContext *app.RequestContext) *DeleteMeetPlansService {
	return &DeleteMeetPlansService{RequestContext: RequestContext, Context: ctx, DAO: query.Plan.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *DeleteMeetPlansService) Run(req *model.DeleteMeetPlansRequest, resp *model.DeleteMeetPlansResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in DeleteMeetPlansService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.DeleteMeetPlansResponse)
	}
	req.Ids = lo.Uniq(req.Ids)
	_, e := h.DAO.Select(query.Plan.Orders.Field()).Where(query.Plan.ID.In(req.Ids...)).Delete()
	if e != nil {
		return errno.ToInternalErr(e)
	}
	return
}
