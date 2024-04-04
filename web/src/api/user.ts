import { http } from "@/utils/http";

export type LoginResult = {
  code: number;
  data: {
    accessToken: string;
    refreshToken: string;
    expires: Date;
  };
  error: string;
};

/** 登录 */
export const login = (code: string) => {
  return http.request<LoginResult>("post", "/api/v1/login", {
    data: {
      code: code
    }
  });
};

/** 刷新token */
export const refreshToken = (refreshToken: string) => {
  return http.request<LoginResult>("post", "/api/v1/login/refresh", {
    data: {
      refreshToken: refreshToken
    }
  });
};

export type UserInfo = {
  id?: string;
  createdAt?: string;
  updatedAt?: string;
  isActive?: boolean;
  isAdmin?: boolean;
  isTeacher?: boolean;
  name?: string;
  pkuID?: string;
  email?: string;
  emailConfirming?: string;
  phoneNumber?: string;
  gender?: string;
  birthday?: string;
  avatar?: string;
  departmentID?: string;
  department?: string;
  office?: string;
  introduction?: string;
  dorm?: string;
  majorID?: string;
  major?: string;
  gradeID?: string;
  grade?: string;
  isGraduated?: boolean;
};

export type GetSelfInfoResult = {
  code: number;
  data?: UserInfo;
  error: string;
};

export const getSelfInfo = () => {
  return http.request<GetSelfInfoResult>("get", "/api/v1/users/self");
};

export const uploadUserInfoAvatarApi = (params?: object, data?: object) => {
  return http.request<GetSelfInfoResult>(
    "put",
    "/api/system/userinfo/self/upload",
    params,
    data
  );
};

export interface PageParams {
  page: number;
  pageSize: number;
}

export interface FilterParams {
  name?: string;
  pkuID?: string;
  isActive?: boolean;
  isTeacher?: boolean;
  isAdmin?: boolean;
  departmentID?: string[];
  majorID?: string[];
  gradeID?: string[];
}

export interface QueryUserParams extends PageParams, FilterParams {}

export interface QueryUserResult {
  code: number;
  data?: UserInfo[];
  error: string;
  pageInfo: {
    page: number;
    pageSize: number;
    total: number;
  };
}

export const searchUser = (param: QueryUserParams) => {
  let params = {
    page: param.page,
    pageSize: param.pageSize
  };
  if (param.name) {
    params["name"] = param.name;
  }
  if (param.pkuID) {
    params["pkuID"] = param.pkuID;
  }
  if (param.isActive !== undefined) {
    params["isActive"] = param.isActive;
  }
  if (param.isTeacher !== undefined) {
    params["isTeacher"] = param.isTeacher;
  }
  if (param.isAdmin !== undefined) {
    params["isAdmin"] = param.isAdmin;
  }
  if (param.departmentID) {
    params["departmentID"] = param.departmentID.join(",");
  }
  if (param.majorID) {
    params["majorID"] = param.majorID.join(",");
  }
  if (param.gradeID) {
    params["gradeID"] = param.gradeID.join(",");
  }

  return http.request<QueryUserResult>("get", "/api/v1/users", {
    params: params
  });
};

export interface CreateUsersResult {
  code: number;
  data?: UserInfo[];
  error: string;
}

export const createUsers = (data: UserInfo[]) => {
  return http.request<CreateUsersResult>("post", `/api/v1/users/`, {
    data: data
  });
};

export const updateUserInfoApi = (data: UserInfo) => {
  return http.request<GetSelfInfoResult>("put", `/api/v1/users/${data.id}`, {
    data: data
  });
};
