include 'base.thrift'

namespace go pkuphy.meetplan.plan

struct Plan {
    1: optional i64 id (vt.ge = "1")
    2: optional i64 teacher_id (vt.ge = "1")
    3: optional i64 start_time (vt.ge = "946684800") // 2000-01-01 00:00:00
    4: optional i64 duration (vt.ge = "1800") // 30*60
    5: optional string place
    6: optional string message
    7: optional i8 quota (vt.in = "1", vt.in = "2")
    8: optional i8 remaining_quota (vt.ge = "0", vt.le = "$quota")
}

struct GetPlanReq{
    1: i64 id
}
struct GetPlanResp{
    1: Plan plan
    255: base.BaseResp base_resp
}

struct MGetPlanReq {
    1: list<i64> id_list (vt.min_size = "1")
}

struct MGetPlanResp {
    1: list<Plan> plan_list

    255: base.BaseResp base_resp
}

struct QueryPlanReq{
    1: optional base.PageParam page_param
    2: optional list<i64> teacher_id_list
    3: optional i64 start_time
}

struct QueryPlanResp{
    1: base.PageParam page_param
    2: list<Plan> plan_list
    255: base.BaseResp base_resp
}

struct CreatePlanReq{
    1: Plan plan (vt.not_nil = "true")
}

struct CreatePlanResp{
    1: Plan plan
    255: base.BaseResp base_resp
}

struct MCreatePlanReq {
    1: list<Plan> plan_list (vt.min_size = "1")
}

struct MCreatePlanResp {
    1: list<Plan> plan_list
    255: base.BaseResp base_resp
}

struct UpdatePlanReq{
    1: Plan plan (vt.not_nil = "true")
}

service Service {
    GetPlanResp GetPlan(1: GetPlanReq req)
    MGetPlanResp MGetPlan(1: MGetPlanReq req)
    QueryPlanResp QueryPlan(1: QueryPlanReq req)

    CreatePlanResp CreatePlan(1: CreatePlanReq req)
    MCreatePlanResp MCreatePlan(1: MCreatePlanReq req)
}