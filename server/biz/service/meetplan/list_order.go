package meetplan

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/biz/dal/pack"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
	"meetplan/pkg/httputil"
)

type ListOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IOrderDo
}

func NewListOrderService(ctx context.Context, RequestContext *app.RequestContext) *ListOrderService {
	return &ListOrderService{RequestContext: RequestContext, Context: ctx, DAO: query.Order.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *ListOrderService) Run(req *model.ListOrderRequest, resp *model.ListOrderResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in ListOrderService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.ListOrderResponse)
	}
	dao := h.DAO
	if len(req.Id) > 0 {
		dao = dao.Where(query.Order.ID.In(req.Id...))
	}
	if len(req.StudentId) > 0 {
		dao = dao.Where(query.Order.StudentID.In(req.StudentId...))
	}
	if len(req.MeetPlanId) > 0 {
		dao = dao.Where(query.Order.PlanID.In(req.MeetPlanId...))
	}
	if req.WithStudent {
		dao = dao.Preload(query.Order.Student)
	}
	if req.WithMeetPlan {
		dao = dao.Preload(query.Order.Plan)
	}
	if req.WithTeacher {
		dao = dao.Preload(query.Order.Teacher)
	}
	offset, limit := httputil.GetPageParam(req.PageParam)
	orders, count, e := dao.FindByPage(offset, limit)
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = pack.OrdersDal2Vo(orders)
	resp.PageParam = &model.Pagination{
		PageNo:     req.PageParam.PageNo,
		PageSize:   req.PageParam.PageSize,
		TotalCount: count,
	}

	return
}
