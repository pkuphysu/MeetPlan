package option

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"go.mongodb.org/mongo-driver/bson"

	"meetplan/api/v1/types"
	"meetplan/model"
	"meetplan/model/query"
)

type DeleteOptionRequest struct {
	OptionID string `path:"id"`
	Name     string `json:"name"`
}

func DeleteOption(ctx context.Context, c *app.RequestContext, req *DeleteOptionRequest) (bool, *types.PageInfo, error) {
	option, err := query.OptionColl.FindByIDStr(ctx, req.OptionID)
	if err != nil {
		return false, nil, err
	}
	if option.Name == model.OptionNameSemesterStartDate || option.Name == model.OptionNameSemesterEndDate {
		return false, nil, errors.New("cannot delete semester start or end date")
	}
	if string(option.Name) != req.Name {
		return false, nil, errors.New("option name does not match")
	}

	_, err = query.OptionColl.Raw().DeleteOne(ctx, bson.M{"_id": option.ID})
	if err != nil {
		return false, nil, err
	}
	return true, nil, nil
}
