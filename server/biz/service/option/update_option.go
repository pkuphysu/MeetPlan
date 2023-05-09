package option

import (
	"context"
	"errors"

	"meetplan/biz/gorm_gen"

	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type UpdateOptionService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IOptionDo
}

func NewUpdateOptionService(ctx context.Context, RequestContext *app.RequestContext) *UpdateOptionService {
	return &UpdateOptionService{RequestContext: RequestContext, Context: ctx, DAO: query.Option.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *UpdateOptionService) Run(req *model.UpdateOptionRequest, resp *model.UpdateOptionResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in UpdateOptionService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.UpdateOptionResponse)
	}

	option, e := h.DAO.Where(query.Option.Name.Eq(req.Key)).First()
	if e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
		return errno.ToInternalErr(e)
	}

	if req.Value == nil {
		if option == nil {
			return
		}
		_, e := h.DAO.Where(query.Option.Name.Eq(req.Key)).Delete()
		if e != nil {
			return errno.ToInternalErr(e)
		}
		return
	} else {
		value := datatypes.JSON{}
		e := value.Scan(*req.Value)
		if e != nil {
			return errno.ToValidationErr(e)
		}
		if option == nil {
			e := h.DAO.Create(&gorm_gen.Option{Name: req.Key, Value: value})
			if e != nil {
				return errno.ToInternalErr(e)
			}
			return
		} else {
			_, e := h.DAO.Where(query.Option.ID.Eq(option.ID)).UpdateSimple(query.Option.Value.Value(value))
			if e != nil {
				return errno.ToInternalErr(e)
			}
			return
		}
	}
}
