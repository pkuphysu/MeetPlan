import { http } from "@/utils/http";

export type Department = {
  id?: string;
  department?: string;
};

export type GetDepartmentListResult = {
  code: number;
  data: Department[];
  error: string;
  pageInfo?: {
    page: number;
    pageSize: number;
    total: number;
  };
};

export const getDepartmentList = (
  page: number,
  pageSize: number,
  search?: string
) => {
  return http.request<GetDepartmentListResult>("get", "/api/v1/departments", {
    params: {
      page: page,
      pageSize: pageSize,
      search: search
    }
  });
};

export type GetDepartmentResult = {
  code: number;
  data?: Department;
  error: string;
};

export const createDepartment = (department: Department) => {
  return http.request<GetDepartmentResult>("post", "/api/v1/departments", {
    data: {
      department: department.department
    }
  });
};

export const updateDepartment = (department: Department) => {
  return http.request<GetDepartmentResult>(
    "put",
    `/api/v1/departments/${department.id}`,
    {
      data: {
        department: department.department
      }
    }
  );
};
