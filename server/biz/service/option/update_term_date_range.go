package option

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/datatypes"

	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/constant"
	"meetplan/pkg/errno"
)

type UpdateTermDateRangeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IOptionDo
}

func NewUpdateTermDateRangeService(ctx context.Context, RequestContext *app.RequestContext) *UpdateTermDateRangeService {
	return &UpdateTermDateRangeService{RequestContext: RequestContext, Context: ctx, DAO: query.Option.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *UpdateTermDateRangeService) Run(req *model.UpdateTermDateRangeRequest, resp *model.UpdateTermDateRangeResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in UpdateTermDateRangeService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.UpdateTermDateRangeResponse)
	}
	record := model.TermDateRange{
		Start: req.Start,
		End:   req.End,
	}
	recordCol := &datatypes.JSON{}
	e := recordCol.Scan(&record)
	if e != nil {
		return errno.ToInternalErr(e)
	}
	option, e := h.DAO.Where(query.Option.Name.Eq(constant.OptionNameTermDateRange)).Assign(query.Option.Value.Value(recordCol)).FirstOrCreate()
	if e != nil {
		return errno.ToInternalErr(e)
	}

	bytes, e := option.Value.MarshalJSON()
	if e != nil {
		return errno.ToInternalErr(e)
	}
	e = json.Unmarshal(bytes, &record)
	resp.Data = &record
	return
}
