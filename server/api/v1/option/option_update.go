package option

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type UpdateOptionRequest struct {
	OptionID string  `path:"id"`
	Name     *string `json:"name"`
	Value    *string `json:"value"`
}

func UpdateOption(ctx context.Context, c *app.RequestContext, req *UpdateOptionRequest) (*model.Option, *types.PageInfo, error) {
	option, err := query.OptionColl.FindByIDStr(ctx, req.OptionID)
	if err != nil {
		return nil, nil, err
	}
	if req.Name != nil {
		option.Name = model.OptionName(*req.Name)
	}
	if req.Value != nil {
		option.Value = *req.Value
	}

	err = query.OptionColl.Upsert(ctx, option)
	if err != nil {
		return nil, nil, err
	}
	return option, nil, nil
}
