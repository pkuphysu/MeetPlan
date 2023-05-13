interface QueryPageParam {
  page_no: number;
  page_size: number;
}

interface Pagination {
  page_no: number;
  page_size: number;
  total_count: string;
}

enum Gender {
  GENDER_MALE = 0,
  GENDER_FEMALE = 1,
}

interface User {
  id: string;
  pku_id: string;
  name: string;
  email: string;
  is_teacher: boolean;
  is_admin: boolean;
  is_active: boolean;
  gender: Gender;
  avatar: string;
  phone: string;
  department: string;
  major: string;
  grade: string;
  dorm: string;
  office: string;
  introduction: string;
  email_change: string;
}

interface LoginRequest {
  code: string;
  nonce?: string;
}

interface LoginResponse {
  code: number;
  message: string;
  data: string;
}

interface GetSelfRequest {
}

interface GetSelfResponse {
  code: number;
  message: string;
  data: User;
}

interface GetUserRequest {
  id: string;
}

interface GetUserResponse {
  code: number;
  message: string;
  data: User;
}

interface ListUserRequest {
  page_param: QueryPageParam;
  is_active: boolean;
  is_teacher: boolean;
  is_admin: boolean;
  ids: string[];
  name: string;
}

interface ListUserResponse {
  code: number;
  message: string;
  page_param: Pagination;
  data: User[];
}

interface CreateUserRequest {
  pku_id: string;
  name: string;
  email: string;
  is_teacher: boolean;
  is_admin: boolean;
  gender: Gender;
  avatar: string;
  phone: string;
  department: string;
  major: string;
  grade: string;
  dorm: string;
  office: string;
  introduction: string;
}

interface CreateUserResponse {
  code: number;
  message: string;
  data: User;
}

interface UpdateUserRequest {
  id: string;
  pku_id: string;
  name: string;
  email: string;
  is_teacher: boolean;
  is_admin: boolean;
  is_active: boolean;
  gender: Gender;
  avatar: string;
  phone: string;
  department: string;
  major: string;
  grade: string;
  dorm: string;
  office: string;
  introduction: string;
}

interface UpdateUserResponse {
  code: number;
  message: string;
  data: User;
}


interface MeetPlan {
  id: string;
  teacher_id: string;
  start_time: string;
  duration: string;
  place: string;
  message: string;
  quota: number;
  teacher: User;
  orders: Order[];
}

enum OrderStatus {
  ORDER_STATUS_CREATED = 0,
  ORDER_STATUS_FINISHED = 1,
  ORDER_STATUS_CANCELLED = 2,
}

interface Order {
  id: string;
  meet_plan_id: string;
  student_id: string;
  message: string;
  status: OrderStatus;
  meet_plan: MeetPlan;
  student: User;
}

interface GetMeetPlanRequest {
  id: string;
  with_teacher: boolean;
  with_orders: boolean;
  with_students: boolean;
}

interface GetMeetPlanResponse {
  code: number;
  message: string;
  data: MeetPlan;
}

interface ListMeetPlanRequest {
  id: string[];
  teacher_id: string[];
  start_time_ge: string;
  student_id: string[];
  with_teacher: boolean;
  with_orders: boolean;
  with_students: boolean;
  page_param: QueryPageParam;
}

interface ListMeetPlanResponse {
  code: number;
  message: string;
  data: MeetPlan[];
  page_param: Pagination;
}

interface CreateMeetPlanRequest {
  teacher_id: string;
  start_time: string;
  duration: string;
  place: string;
  message: string;
  quota: number;
}

interface CreateMeetPlanResponse {
  code: number;
  message: string;
  data: MeetPlan;
}

interface UpdateMeetPlanRequest {
  id: string;
  start_time: string;
  duration: string;
  place: string;
  message: string;
  quota: number;
}

interface UpdateMeetPlanResponse {
  code: number;
  message: string;
  data: MeetPlan;
}

interface DeleteMeetPlanRequest {
  id: string;
}

interface DeleteMeetPlanResponse {
  code: number;
  message: string;
}

interface DeleteMeetPlansRequest {
  ids: string[];
}

interface DeleteMeetPlansResponse {
  code: number;
  message: string;
}

interface GetOrderRequest {
  id: string;
  with_meet_plan: boolean;
  with_student: boolean;
  with_teacher: boolean;
}

interface GetOrderResponse {
  code: number;
  message: string;
  data: Order;
}

interface ListOrderRequest {
  id: string[];
  meet_plan_id: string[];
  student_id: string[];
  status: OrderStatus;
  with_meet_plan: boolean;
  with_student: boolean;
  with_teacher: boolean;
  page_param: QueryPageParam;
}

interface ListOrderResponse {
  code: number;
  message: string;
  data: Order[];
  page_param: Pagination;
}

interface CreateOrderRequest {
  meet_plan_id: string;
  student_id: string;
  message: string;
}

interface CreateOrderResponse {
  code: number;
  message: string;
  data: Order;
}

interface UpdateOrderRequest {
  id: string;
  message: string;
  status: OrderStatus;
}

interface UpdateOrderResponse {
  code: number;
  message: string;
  data: Order;
}

interface CreateMeetPlanAndOrderRequest {
  teacher_id: string;
  start_time: string;
  duration: string;
  place: string;
  message: string;
  quota: number;
  orders: Order[];
}

interface Order {
  student_id: string;
  message: string;
  status: OrderStatus;
}

interface CreateMeetPlanAndOrderResponse {
  code: number;
  message: string;
  data: MeetPlan;
}


interface FriendLink {
  name: string;
  url: string;
  description: string;
}

interface ListFriendLinkRequest {
}

interface ListFriendLinkResponse {
  code: number;
  message: string;
  data: FriendLink[];
}

interface CreateFriendLinkRequest {
  name: string;
  url: string;
  description: string;
}

interface CreateFriendLinkResponse {
  code: number;
  message: string;
  data: FriendLink;
}

interface UpdateRecord {
  timestamp: string;
  author: string;
  url: string;
  description: string;
}

interface ListUpdateRecordRequest {
}

interface ListUpdateRecordResponse {
  code: number;
  message: string;
  data: UpdateRecord[];
}

interface CreateUpdateRecordRequest {
  author: string;
  url: string;
  description: string;
  timestamp: string;
}

interface CreateUpdateRecordResponse {
  code: number;
  message: string;
  data: UpdateRecord;
}

interface TermDateRange {
  start: string;
  end: string;
}

interface GetTermDateRangeRequest {
}

interface GetTermDateRangeResponse {
  code: number;
  message: string;
  data: TermDateRange;
}

interface UpdateTermDateRangeRequest {
  start: string;
  end: string;
}

interface UpdateTermDateRangeResponse {
  code: number;
  message: string;
  data: TermDateRange;
}

interface GetOptionRequest {
  key: string;
}

interface GetOptionResponse {
  code: number;
  message: string;
  data: string;
}

interface UpdateOptionRequest {
  key: string;
  value: string;
}

interface UpdateOptionResponse {
  code: number;
  message: string;
}

