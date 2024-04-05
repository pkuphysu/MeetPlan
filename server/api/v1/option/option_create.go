package option

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type CreateOptionRequest struct {
	Options []*model.Option `json:"options"`
}

func CreateOptions(ctx context.Context, c *app.RequestContext, req *CreateOptionRequest) ([]*model.Option, *types.PageInfo, error) {
	for _, option := range req.Options {
		option.ID = primitive.NewObjectID()
	}
	err := query.OptionColl.InsertMany(ctx, req.Options)
	if err != nil {
		return nil, nil, err
	}
	return req.Options, nil, nil
}
