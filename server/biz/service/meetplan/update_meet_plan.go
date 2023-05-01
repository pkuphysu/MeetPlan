package meetplan

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gen/field"
	"gorm.io/gorm"

	"meetplan/biz/dal/pack"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type UpdateMeetPlanService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IPlanDo
}

func NewUpdateMeetPlanService(ctx context.Context, RequestContext *app.RequestContext) *UpdateMeetPlanService {
	return &UpdateMeetPlanService{RequestContext: RequestContext, Context: ctx, DAO: query.Plan.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *UpdateMeetPlanService) Run(req *model.UpdateMeetPlanRequest, resp *model.UpdateMeetPlanResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in UpdateMeetPlanService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.UpdateMeetPlanResponse)
	}

	var fields []field.AssignExpr
	if req.StartTime != nil {
		fields = append(fields, query.Plan.StartTime.Value(time.Unix(*req.StartTime, 0)))
	}
	if req.Duration != nil {
		fields = append(fields, query.Plan.Duration.Value(*req.Duration))
	}
	if req.Place != nil {
		fields = append(fields, query.Plan.Place.Value(*req.Place))
	}
	if req.Message != nil {
		fields = append(fields, query.Plan.Message.Value(*req.Message))
	}
	if req.Quota != nil {
		fields = append(fields, query.Plan.Quota.Value(int8(*req.Quota)))
	}
	if len(fields) == 0 {
		return errno.NewValidationErr("no field to update")
	}

	_, e := h.DAO.Where(query.Plan.ID.Eq(req.Id)).UpdateColumnSimple(fields...)
	if e != nil && errors.Is(e, gorm.ErrRecordNotFound) {
		return errno.NewNotFoundErr("plan not found")
	} else if e != nil {
		return errno.ToInternalErr(e)
	}

	plan, e := h.DAO.Where(query.Plan.ID.Eq(req.Id)).First()
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = pack.PlanDal2Vo(plan)
	return
}
