package meetplan

import (
	"context"
	"time"

	"gorm.io/gen/field"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/biz/dal/pack"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
	"meetplan/pkg/httputil"
)

type ListMeetPlanService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	PlanDAO        query.IPlanDo
}

func NewListMeetPlanService(ctx context.Context, RequestContext *app.RequestContext) *ListMeetPlanService {
	return &ListMeetPlanService{RequestContext: RequestContext, Context: ctx, PlanDAO: query.Plan.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *ListMeetPlanService) Run(req *model.ListMeetPlanRequest, resp *model.ListMeetPlanResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in ListMeetPlanService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.ListMeetPlanResponse)
	}
	dao := h.PlanDAO.Where(query.Plan.StartTime.Gte(time.Unix(req.StartTimeGe, 0)))
	if len(req.Id) > 0 {
		dao = dao.Where(query.Plan.ID.In(req.Id...))
	}
	if len(req.TeacherId) > 0 {
		dao = dao.Where(query.Plan.TeacherID.In(req.TeacherId...))
	}
	if len(req.StudentId) > 0 {
		dao = dao.Order(query.Order.StudentID.In(req.StudentId...))
	}
	if req.WithTeacher {
		dao = dao.Preload(query.Plan.Teacher)
	}
	if req.WithOrders || req.WithStudents {
		dao = dao.Preload(query.Plan.Orders)
		if req.WithStudents {
			dao = dao.Preload(field.NewRelation("Orders.Student", ""))
		}
	}
	offset, limit, param := httputil.GetPageParam(req.PageParam)
	res, count, e := dao.Order(query.Plan.ID.Desc()).FindByPage(offset, limit)
	if e != nil {
		return errno.NewInternalErr(e.Error())
	}
	resp.Data = pack.PlansDal2Vo(res)
	resp.PageParam = &model.Pagination{
		PageNo:     param.PageNo,
		PageSize:   param.PageSize,
		TotalCount: count,
	}

	return
}
