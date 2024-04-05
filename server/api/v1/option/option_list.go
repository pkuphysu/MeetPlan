package option

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type ListOptionRequest struct {
	Search   string `query:"search"`
	Name     string `query:"name"`
	Page     int    `query:"page"`
	PageSize int    `query:"pageSize"`
}

func ListOption(ctx context.Context, c *app.RequestContext, req *ListOptionRequest) ([]*model.Option, *types.PageInfo, error) {
	filter := bson.M{}
	if req.Search != "" {
		filter["name"] = bson.M{"$regex": req.Search}
	}
	if req.Name != "" {
		filter["name"] = strings.Split(req.Name, ",")
	}
	options, err := query.OptionColl.FindPage(ctx, filter, req.Page, req.PageSize)
	if err != nil {
		return nil, nil, err
	}
	total, err := query.OptionColl.Count(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	pageInfo := &types.PageInfo{
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	return options, pageInfo, nil
}
