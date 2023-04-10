include "base.thrift"

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

enum OrderStatus {
    CREATED = 1
    FINISHED = 2
    CANCELLED = 3
}

struct Order {
    1: optional i64 id (vt.ge = "1")
    2: optional i64 plan_id (vt.ge = "1")
    3: optional i64 student_id (vt.ge = "1")
    4: optional string message
    5: optional OrderStatus status (vt.defined_only = "true")
}

struct GetOrderReq {
    1: i64 id (vt.ge = "1")
}

struct GetOrderResp {
    1: Order order
    255: base.BaseResp base_resp
}

struct MGetOrderReq {
    1: list<i64> ids (vt.min_size = "1")
}

struct MGetOrderResp {
    1: list<Order> orders
    255: base.BaseResp base_resp
}

struct QueryOrderReq {
    1: optional base.PageParam page_param
    2: optional list<i64> plan_ids (vt.min_size = "1")
    3: optional list<i64> student_ids (vt.min_size = "1")
    4: optional OrderStatus status (vt.defined_only = "true")
    5: optional list<i64> teacher_ids (vt.min_size = "1")
}

struct QueryOrderResp {
    1: base.PageParam page_param
    2: list<Order> orders
    255: base.BaseResp base_resp
}

struct CreateOrderReq {
    1: i64 plan_id (vt.ge = "1")
    2: i64 student_id (vt.ge = "1")
    3: optional string message
    4: optional OrderStatus status (vt.defined_only = "true")
}

struct CreateOrderResp {
    1: Order order
    255: base.BaseResp base_resp
}

struct MCreateOrderReq {
    1: list<CreateOrderReq> orders (vt.min_size = "1")
}

struct MCreateOrderResp {
    1: list<Order> orders
    255: base.BaseResp base_resp
}

struct UpdateOrderReq {
    1: i64 id (vt.ge = "1")
    2: optional string message
    3: optional OrderStatus status (vt.defined_only = "true")
}

struct UpdateOrderResp {
    1: Order order
    255: base.BaseResp base_resp
}

service Service {
    GetPlanResp GetPlan(1: GetPlanReq req)
    MGetPlanResp MGetPlan(1: MGetPlanReq req)
    QueryPlanResp QueryPlan(1: QueryPlanReq req)

    CreatePlanResp CreatePlan(1: CreatePlanReq req)
    MCreatePlanResp MCreatePlan(1: MCreatePlanReq req)

    GetOrderResp GetOrder(1: GetOrderReq req)
    MGetOrderResp MGetOrder(1: MGetOrderReq req)
    QueryOrderResp QueryOrder(1: QueryOrderReq req)

    CreateOrderResp CreateOrder(1: CreateOrderReq req)
    MCreateOrderResp MCreateOrder(1: MCreateOrderReq req)

    UpdateOrderResp UpdateOrder(1: UpdateOrderReq req)
}