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

type CreateMeetPlanService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IPlanDo
	UserDAO        query.IUserDo
}

func NewCreateMeetPlanService(ctx context.Context, RequestContext *app.RequestContext) *CreateMeetPlanService {
	return &CreateMeetPlanService{RequestContext: RequestContext, Context: ctx, DAO: query.Plan.WithContext(ctx), UserDAO: query.User.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *CreateMeetPlanService) Run(req *model.CreateMeetPlanRequest, resp *model.CreateMeetPlanResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateMeetPlanService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.CreateMeetPlanResponse)
	}

	teacher, e := h.UserDAO.Where(query.User.ID.Eq(req.TeacherId)).First()
	if e != nil && errors.Is(e, gorm.ErrRecordNotFound) {
		return errno.NewNotFoundErr("teacher not found")
	} else if e != nil {
		return errno.ToInternalErr(e)
	}

	office := ""
	if teacher.Office != nil {
		office = *teacher.Office
	}

	plan := &gorm_gen.Plan{
		TeacherID: req.TeacherId,
		StartTime: time.Unix(req.StartTime, 0),
		Duration:  req.Duration,
		Place:     lo.If(req.Place != "", req.Place).Else(office),
		Quota:     int8(req.Quota),
		Message:   &req.Message,
	}
	if len(plan.Place) == 0 {
		return errno.NewValidationErr("plan place can not be empty")
	}
	e = h.DAO.Create(plan)
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = &model.MeetPlan{
		Id:        plan.ID,
		TeacherId: plan.TeacherID,
		StartTime: plan.StartTime.Unix(),
		Duration:  plan.Duration,
		Place:     plan.Place,
		Message: lo.IfF(plan.Message != nil, func() string {
			return *plan.Message
		}).Else(""),
		Quota:   int32(plan.Quota),
		Teacher: pack.UserDal2Vo(teacher),
		Orders:  nil,
	}
	return
}
