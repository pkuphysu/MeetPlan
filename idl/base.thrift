namespace go pkuphy.meetplan.base

enum StatusCode {
    SuccessCode = 0
    ServiceErrCode = 10001
    ParamErrCode = 10002
    AuthorizationFailedErrCode = 10003
    // User
    UserNotFoundErrCode = 10004
    UserCannotLoginErrCode = 10005
}

struct BaseResp{
    1: StatusCode status_code
    2: string message
}

struct PageParam{
    1: i32 page_num (vt.ge = "1")
    2: i32 page_size (vt.ge = "10", vt.le = "50")
}
