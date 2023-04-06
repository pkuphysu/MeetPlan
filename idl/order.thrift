include 'base.thrift'

namespace go pkuphy.meetplan.order

enum OrderStatus {
    CREATED = 1
    FINISHED = 2
    Canceled = 3
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
    3: optional list<i64> student_id (vt.min_size = "1")
    4: optional OrderStatus status (vt.defined_only = "true")
}

struct QueryOrderResp {
    1: base.PageParam page_result
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
    GetOrderResp GetOrder(1: GetOrderReq req)
    MGetOrderResp MGetOrder(1: MGetOrderReq req)
    QueryOrderResp QueryOrder(1: QueryOrderReq req)

    CreateOrderResp CreateOrder(1: CreateOrderReq req)
    MCreateOrderResp MCreateOrder(1: MCreateOrderReq req)

    UpdateOrderResp UpdateOrder(1: UpdateOrderReq req)
}

