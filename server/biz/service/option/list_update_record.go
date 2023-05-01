package option

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type ListUpdateRecordService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IUpdateRecordDo
}

func NewListUpdateRecordService(ctx context.Context, RequestContext *app.RequestContext) *ListUpdateRecordService {
	return &ListUpdateRecordService{RequestContext: RequestContext, Context: ctx, DAO: query.UpdateRecord.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *ListUpdateRecordService) Run(req *model.ListUpdateRecordRequest, resp *model.ListUpdateRecordResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in ListUpdateRecordService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.ListUpdateRecordResponse)
	}
	records, e := h.DAO.Find()
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = make([]*model.UpdateRecord, len(records))
	for i, record := range records {
		resp.Data[i] = &model.UpdateRecord{
			Timestamp:   record.Time.Unix(),
			Author:      record.Author,
			Url:         record.URL,
			Description: record.Info,
		}
	}
	return
}
