package option

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/datatypes"

	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/constant"
	"meetplan/pkg/errno"
)

type GetTermDateRangeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IOptionDo
}

func NewGetTermDateRangeService(ctx context.Context, RequestContext *app.RequestContext) *GetTermDateRangeService {
	return &GetTermDateRangeService{RequestContext: RequestContext, Context: ctx, DAO: query.Option.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *GetTermDateRangeService) Run(req *model.GetTermDateRangeRequest, resp *model.GetTermDateRangeResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in GetTermDateRangeService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.GetTermDateRangeResponse)
	}

	record := model.TermDateRange{
		Start: time.Now().Unix(),
		End:   time.Now().Add(time.Hour * 24 * 30 * 6).Unix(),
	}
	recordCol := &datatypes.JSON{}
	bytes, e := json.Marshal(&record)
	if e != nil {
		return errno.ToInternalErr(e)
	}
	e = recordCol.Scan(bytes)
	if e != nil {
		return errno.ToInternalErr(e)
	}

	option, e := h.DAO.Where(query.Option.Name.Eq(constant.OptionNameTermDateRange)).Attrs(query.Option.Value.Value(recordCol)).FirstOrInit()
	if e != nil {
		return errno.ToInternalErr(e)
	}

	bytes, e = option.Value.MarshalJSON()
	if e != nil {
		return errno.ToInternalErr(e)
	}
	e = json.Unmarshal(bytes, &record)
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = &record
	return
}
