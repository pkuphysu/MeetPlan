package option

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type CreateOptionRequest struct {
	Options []*model.Option `json:"options"`
}

func CreateOptions(ctx context.Context, c *app.RequestContext, req *CreateOptionRequest) ([]*model.Option, *types.PageInfo, error) {
	err := query.OptionColl.UpsertMany(ctx, req.Options)
	if err != nil {
		return nil, nil, err
	}
	return req.Options, nil, nil
}
