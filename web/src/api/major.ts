import { http } from "@/utils/http";

export type Major = {
  id?: string;
  major?: string;
};

export type GetMajorListResult = {
  code: number;
  data: Major[];
  error: string;
  pageInfo?: {
    page: number;
    pageSize: number;
    total: number;
  };
};

export const getMajorList = (
  page: number,
  pageSize: number,
  search?: string
) => {
  return http.request<GetMajorListResult>("get", "/api/v1/majors", {
    params: {
      page: page,
      pageSize: pageSize,
      search: search
    }
  });
};

export type GetMajorResult = {
  code: number;
  data?: Major;
  error: string;
};

export const createMajor = (major: Major) => {
  return http.request<GetMajorResult>("post", "/api/v1/majors", {
    data: {
      major: major.major
    }
  });
};

export const updateMajor = (major: Major) => {
  return http.request<GetMajorResult>("put", `/api/v1/majors/${major.id}`, {
    data: {
      major: major.major
    }
  });
};
