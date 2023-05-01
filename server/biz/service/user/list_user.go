package user

import (
	"context"

	"meetplan/biz/dal/pack"
	"meetplan/pkg/httputil"

	"meetplan/biz/gorm_gen/query"

	"github.com/cloudwego/hertz/pkg/app"

	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type ListUserService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	UserDAO        query.IUserDo
}

func NewListUserService(ctx context.Context, RequestContext *app.RequestContext) *ListUserService {
	return &ListUserService{RequestContext: RequestContext, Context: ctx, UserDAO: query.User.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *ListUserService) Run(req *model.ListUserRequest, resp *model.ListUserResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in ListUserService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.ListUserResponse)
	}

	dao := h.UserDAO
	if len(req.Ids) > 0 {
		dao = dao.Where(query.User.ID.In(req.Ids...))
	}
	if req.IsTeacher != nil {
		dao = dao.Where(query.User.IsTeacher.Is(*req.IsTeacher))
	}
	if req.IsActive != nil {
		dao = dao.Where(query.User.IsActive.Is(*req.IsActive))
	}
	if req.IsAdmin != nil {
		dao = dao.Where(query.User.IsAdmin.Is(*req.IsAdmin))
	}
	offset, limit := httputil.GetPageParam(req.PageParam)
	users, count, e := dao.FindByPage(offset, limit)
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = make([]*model.User, 0, len(users))
	for _, u := range users {
		resp.Data = append(resp.Data, pack.UserDal2Vo(u))
	}
	resp.PageParam = &model.Pagination{
		PageNo:     req.PageParam.PageNo,
		PageSize:   req.PageParam.PageSize,
		TotalCount: count,
	}

	return
}
