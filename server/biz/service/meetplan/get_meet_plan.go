package meetplan

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gen/field"
	"gorm.io/gorm"

	"meetplan/biz/dal/pack"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type GetMeetPlanService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	PlanDAO        query.IPlanDo
}

func NewGetMeetPlanService(ctx context.Context, RequestContext *app.RequestContext) *GetMeetPlanService {
	return &GetMeetPlanService{RequestContext: RequestContext, Context: ctx, PlanDAO: query.Plan.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *GetMeetPlanService) Run(req *model.GetMeetPlanRequest, resp *model.GetMeetPlanResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in GetMeetPlanService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.GetMeetPlanResponse)
	}
	var preloads []field.RelationField
	if req.WithTeacher {
		preloads = append(preloads, query.Plan.Teacher)
	}
	if req.WithOrders || req.WithStudents {
		preloads = append(preloads, query.Plan.Orders)
		if req.WithStudents {
			preloads = append(preloads, field.NewRelation("Orders.Student", ""))
		}
	}
	plan, e := h.PlanDAO.Where(query.Plan.ID.Eq(req.Id)).Preload(preloads...).First()
	if e != nil && errors.Is(e, gorm.ErrRecordNotFound) {
		return errno.NewNotFoundErr("plan not found")
	} else if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = pack.PlanDal2Vo(plan)
	return
}
