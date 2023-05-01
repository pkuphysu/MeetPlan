package option

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/biz/gorm_gen"
	"meetplan/biz/gorm_gen/query"
	model "meetplan/biz/model"
	"meetplan/pkg/errno"
)

type CreateUpdateRecordService struct {
	RequestContext *app.RequestContext
	Context        context.Context
	DAO            query.IUpdateRecordDo
}

func NewCreateUpdateRecordService(ctx context.Context, RequestContext *app.RequestContext) *CreateUpdateRecordService {
	return &CreateUpdateRecordService{RequestContext: RequestContext, Context: ctx, DAO: query.UpdateRecord.WithContext(ctx)}
}

// Run req should not be nil and resp should not be nil
func (h *CreateUpdateRecordService) Run(req *model.CreateUpdateRecordRequest, resp *model.CreateUpdateRecordResponse) (err *errno.Err) {
	defer func() {
		if e := recover(); e != nil {
			err = errno.NewInternalErr("panic in CreateUpdateRecordService.Run")
		}
	}()
	if resp == nil {
		resp = new(model.CreateUpdateRecordResponse)
	}
	record := &gorm_gen.UpdateRecord{
		Time:   time.Unix(req.Timestamp, 0),
		Author: req.Author,
		URL:    req.Url,
		Info:   req.Description,
	}
	e := h.DAO.Create(record)
	if e != nil {
		return errno.ToInternalErr(e)
	}
	resp.Data = &model.UpdateRecord{
		Timestamp:   record.Time.Unix(),
		Author:      record.Author,
		Url:         record.URL,
		Description: record.Info,
	}
	return
}
