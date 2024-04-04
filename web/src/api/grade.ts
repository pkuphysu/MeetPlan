import { http } from "@/utils/http";

export type Grade = {
  id?: string;
  grade?: string;
  isGraduated?: boolean;
};

export type GetGradeListResult = {
  code: number;
  data: Grade[];
  error: string;
  pageInfo?: {
    page: number;
    pageSize: number;
    total: number;
  };
};

export const getGradeList = (
  page: number,
  pageSize: number,
  search?: string
) => {
  return http.request<GetGradeListResult>("get", "/api/v1/grades", {
    params: {
      page: page,
      pageSize: pageSize,
      search: search
    }
  });
};

export type GetGradeResult = {
  code: number;
  data?: Grade;
  error: string;
};

export const createGrade = (grade: Grade) => {
  return http.request<GetGradeResult>("post", "/api/v1/grades", {
    data: {
      grade: grade.grade,
      isGraduated: grade.isGraduated
    }
  });
};

export const updateGrade = (grade: Grade) => {
  return http.request<GetGradeResult>("put", `/api/v1/grades/${grade.id}`, {
    data: {
      grade: grade.grade,
      isGraduated: grade.isGraduated
    }
  });
};
