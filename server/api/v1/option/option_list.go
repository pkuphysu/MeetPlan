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
	Search string `query:"search"`
	Name   string `query:"name"`
}

func ListOption(ctx context.Context, c *app.RequestContext, req *ListOptionRequest) ([]*model.Option, *types.PageInfo, error) {
	filter := bson.M{}
	if req.Search != "" {
		filter["name"] = bson.M{"$regex": req.Search}
	}
	if req.Name != "" {
		filter["name"] = strings.Split(req.Name, ",")
	}
	options, err := query.OptionColl.FindAll(ctx, filter)
	if err != nil {
		return nil, nil, err
	}
	return options, nil, nil
}
