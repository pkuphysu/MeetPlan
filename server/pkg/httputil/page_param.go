package httputil

import "meetplan/biz/model"

func GetPageParam(pageParam *model.QueryPageParam) (offset, limit int, param *model.QueryPageParam) {
	if pageParam == nil {
		pageParam = &model.QueryPageParam{
			PageNo:   1,
			PageSize: 10,
		}
	}
	param = pageParam

	if pageParam.PageNo < 1 {
		pageParam.PageNo = 1
	}
	if pageParam.PageSize < 1 {
		pageParam.PageSize = 10
	}
	if pageParam.PageSize > 100 {
		pageParam.PageSize = 100
	}
	offset = int((pageParam.PageNo - 1) * pageParam.PageSize)
	limit = int(pageParam.PageSize)
	return
}
