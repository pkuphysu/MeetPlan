package option

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"

	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type GetOptionService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IOptionDo
}

func NewGetOptionService(ctx context.Context, RequestContext *app.RequestContext) *GetOptionService {
	return &GetOptionService{RequestContext: RequestContext, Context: ctx, DAO: query.Option.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *GetOptionService) Run(req *model.GetOptionRequest, resp *model.GetOptionResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in GetOptionService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.GetOptionResponse)
	}
	option, e := h.DAO.Where(query.Option.Name.Eq(req.Key)).First()
	if e != nil && errors.Is(e, gorm.ErrRecordNotFound) {
		return errno.NewNotFoundErr("option not found")
	} else if e != nil {
		return errno.ToInternalErr(e)
	}
	value, e := json.Marshal(option.Value)
	resp.Data = string(value)
	return
}
