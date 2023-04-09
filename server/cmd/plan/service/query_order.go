package service

import (
	"context"
	"github.com/pkuphysu/meetplan/gorm_gen/query"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/base"
	"github.com/pkuphysu/meetplan/kitex_gen/pkuphy/meetplan/plan"
)

type QueryOrderServiceI interface {
	QueryOrder(req *plan.QueryOrderReq) ([]*plan.Order, *base.PageParam, error)
}

func NewQueryOrderService(ctx context.Context) QueryOrderServiceI {
	return &queryOrderService{
		ctx:  ctx,
		dao:  query.Q.WithContext(ctx).Order,
		plan: query.Q.WithContext(ctx).Plan,
	}
}

type queryOrderService struct {
	ctx  context.Context
	dao  query.IOrderDo
	plan query.IPlanDo
}

func (s *queryOrderService) QueryOrder(req *plan.QueryOrderReq) ([]*plan.Order, *base.PageParam, error) {
	dao := s.dao

	if len(req.PlanIds) > 0 {
		dao = dao.Where(query.Q.Order.PlanID.In(req.PlanIds...))
	}
	if len(req.StudentIds) > 0 {
		dao = dao.Where(query.Q.Order.StudentID.In(req.StudentIds...))
	}
	if len(req.TeacherIds) > 0 {
		dao = dao.LeftJoin(s.plan, query.Q.Plan.ID.EqCol(query.Q.Order.PlanID))
		dao = dao.Where(query.Q.Plan.TeacherID.In(req.TeacherIds...))
	}
	if req.Status != nil {
		dao = dao.Where(query.Q.Order.Status.Eq(int8(*req.Status)))
	}
	var pageParam *base.PageParam
	if req.PageParam != nil {
		pageParam = req.PageParam
	} else {
		pageParam = base.NewPageParam()
		pageParam.SetPageNum(1)
		pageParam.SetPageSize(10)
	}
	dao = dao.Limit(int(pageParam.PageSize)).Offset(int(pageParam.PageSize * (pageParam.PageNum - 1)))

	plans, err := dao.Find()
	if err != nil {
		return nil, nil, err
	}

	return packOrders(plans), pageParam, nil
}
