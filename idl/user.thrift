namespace go user

include 'base.thrift'

enum Gender {
    Male = 0
    Female = 1
}

struct User{
    1: optional i64 id (vt.ge = "1")
    2: optional string pku_id (vt.min_size = "10", vt.max_size = "10", vt.pattern = "^[0-9]+$")
    3: optional string name
    4: optional string email
    5: optional bool is_active
    6: optional bool is_teacher
    7: optional bool is_admin
    8: optional Gender gender
    9: optional string avatar
    10: optional string department
    11: optional string phone
    // only for student
    12: optional string major
    13: optional i8 grade
    14: optional string dorm
    // only for teacher
    15: optional string office
    16: optional string introduce
}

struct LoginReq {
    1: required string oauth_code
    2: optional string nonce
}

struct LoginResp {
    1: i64 id
    2: string jwt
    255: base.BaseResp base_resp
}

struct GetUserReq{
    1: optional i64 id (vt.ge = "1")
    2: optional string pku_id (vt.min_size = "10", vt.max_size = "10", vt.pattern = "^[0-9]+$")
}

struct GetUserResp{
    1: User user
    255: base.BaseResp base_resp
}

struct MGetUserReq{
    1: optional base.PageParam page_param
    2: optional bool is_active
    3: optional bool is_teacher
    4: optional bool is_admin
}

struct MGetUserResp{
    1: base.PageParam page_param
    2: list<User> users
    255: base.BaseResp base_resp
}

struct UpdateUserReq{
    1: User user (vt.not_nil = "true")
}

struct UpdateUserResp{
    255: base.BaseResp base_resp
}


service UserService {
    LoginResp Login(1: LoginReq req)

    GetUserResp GetUser(1: GetUserReq req)
    MGetUserResp MGetUser(1: MGetUserReq req)
    UpdateUserResp UpdateUser(1: UpdateUserReq req)
}