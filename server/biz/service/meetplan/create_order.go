package meetplan

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/biz/gorm_gen"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type CreateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IOrderDo
	UserDAO        query.IUserDo
}

func NewCreateOrderService(ctx context.Context, RequestContext *app.RequestContext) *CreateOrderService {
	return &CreateOrderService{RequestContext: RequestContext, Context: ctx,
		DAO: query.Order.WithContext(ctx), UserDAO: query.User.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *CreateOrderService) Run(req *model.CreateOrderRequest, resp *model.CreateOrderResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateOrderService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.CreateOrderResponse)
	}
	student, e := h.UserDAO.Where(query.User.ID.Eq(req.StudentId)).First()
	if e != nil && errors.Is(e, gorm.ErrRecordNotFound) {
		return errno.NewNotFoundErr("student not found")
	} else if e != nil {
		return errno.ToInternalErr(e)
	}
	if student.IsTeacher {
		return errno.NewValidationErr("student_id is not a student")
	}
	e = h.DAO.Create(&gorm_gen.Order{
		PlanID:    req.MeetPlanId,
		StudentID: req.StudentId,
		Message:   &req.Message,
	})
	if e != nil {
		return
	}
	return
}
