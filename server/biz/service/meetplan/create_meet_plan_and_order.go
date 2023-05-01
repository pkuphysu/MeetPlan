package meetplan

import (
	"context"
	"errors"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/samber/lo"
	"gorm.io/gorm"

	"meetplan/biz/dal/pack"
	"meetplan/biz/gorm_gen"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type CreateMeetPlanAndOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	PlanDAO        query.IPlanDo
	UserDAO        query.IUserDo
}

func NewCreateMeetPlanAndOrderService(ctx context.Context, RequestContext *app.RequestContext) *CreateMeetPlanAndOrderService {
	return &CreateMeetPlanAndOrderService{RequestContext: RequestContext, Context: ctx,
		PlanDAO: query.Plan.WithContext(ctx), UserDAO: query.User.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *CreateMeetPlanAndOrderService) Run(req *model.CreateMeetPlanAndOrderRequest, resp *model.CreateMeetPlanAndOrderResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateMeetPlanAndOrderService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.CreateMeetPlanAndOrderResponse)
	}

	// check teacher
	teacher, e := h.UserDAO.Where(query.User.ID.Eq(req.TeacherId)).First()
	if e != nil && errors.Is(e, gorm.ErrRecordNotFound) {
		return errno.NewNotFoundErr("teacher not found")
	} else if e != nil {
		return errno.ToInternalErr(e)
	}

	plan := &gorm_gen.Plan{
		ID:        0,
		TeacherID: req.TeacherId,
		StartTime: time.Unix(req.StartTime, 0),
		Duration:  req.Duration,
		Place: lo.If(req.Place != "", req.Place).ElseIfF(teacher.Office != nil, func() string {
			return *teacher.Office
		}).Else(""),
		Quota:   int8(req.Quota),
		Message: nil,
		Orders:  pack.OrdersVo2Dal(req.Orders),
		Teacher: nil,
	}

	e = h.PlanDAO.Create(plan)
	if e != nil {
		return errno.ToInternalErr(e)
	}

	resp.Data = pack.PlanDal2Vo(plan)
	return
}
