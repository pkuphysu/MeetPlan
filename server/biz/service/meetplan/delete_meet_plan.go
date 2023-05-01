package meetplan

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"meetplan/biz/gorm_gen"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type DeleteMeetPlanService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IPlanDo
}

func NewDeleteMeetPlanService(ctx context.Context, RequestContext *app.RequestContext) *DeleteMeetPlanService {
	return &DeleteMeetPlanService{RequestContext: RequestContext, Context: ctx, DAO: query.Plan.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *DeleteMeetPlanService) Run(req *model.DeleteMeetPlanRequest, resp *model.DeleteMeetPlanResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in DeleteMeetPlanService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.DeleteMeetPlanResponse)
	}

	res, e := h.DAO.Select(query.Plan.Orders.Field()).Delete(&gorm_gen.Plan{ID: req.Id})
	if e != nil {
		return errno.ToInternalErr(e)
	}
	if res.RowsAffected == 0 {
		return errno.NewNotFoundErr("plan not found")
	}
	return
}
