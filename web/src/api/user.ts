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
  return http.request<LoginResult>("post", "/api/v1/login", {
    data: {
      refreshToken: refreshToken
    }
  });
};

export type UserInfo = {
  id: string;
  createdAt: string;
  updatedAt: string;
  isActive: boolean;
  isAdmin: boolean;
  isTeacher: boolean;
  name: string;
  pkuID: string;
  email: string;
  emailConfirming: string;
  phoneNumber: string;
  gender: string;
  birthday: string;
  avatar: string;
  departmentID: string;
  department: string;
  office: string;
  introduction: string;
  dorm: string;
  majorID: string;
  major: string;
  gradeID: string;
  grade: string;
  isGraduated: boolean;
};

export type GetSelfInfoResult = {
  code: number;
  data: UserInfo;
  error: string;
};

export const getSelfInfo = () => {
  return http.request<GetSelfInfoResult>("get", "/api/v1/users/self");
};
