package pack

import (
	"github.com/samber/lo"

	"meetplan/biz/gorm_gen"
	"meetplan/biz/model"
)

func PlanDal2Vo(plan *gorm_gen.Plan) *model.MeetPlan {
	if plan == nil {
		return nil
	}
	return &model.MeetPlan{
		Id:        plan.ID,
		TeacherId: plan.TeacherID,
		StartTime: plan.StartTime.Unix(),
		Duration:  plan.Duration,
		Place:     plan.Place,
		Message: lo.IfF(plan.Message != nil, func() string {
			return *plan.Message
		}).Else(""),
		Quota:   int32(plan.Quota),
		Teacher: UserDal2Vo(plan.Teacher),
		Orders:  OrdersDal2Vo(plan.Orders),
	}
}

func PlansDal2Vo(plans []*gorm_gen.Plan) []*model.MeetPlan {
	if plans == nil {
		return nil
	}
	var res []*model.MeetPlan
	for _, plan := range plans {
		res = append(res, PlanDal2Vo(plan))
	}
	return res
}
