import { http } from "@/utils/http";

export interface Option {
  id?: string;
  name?: string;
  value?: string;
}

export interface OptionResult {
  code: number;
  data?: Option;
  error: string;
}

export interface OptionsResult {
  code: number;
  data?: Option[];
  error: string;
  pageInfo: {
    total: number;
    page: number;
    pageSize: number;
  };
}

export const searchOptions = (
  page: number,
  pageSize: number,
  search?: string
) => {
  return http.request<OptionsResult>("get", "/api/v1/options", {
    params: {
      search: search,
      page: page,
      pageSize: pageSize
    }
  });
};

export const getOption = (page: number, pageSize: number, names: string[]) => {
  return http.request<OptionsResult>("get", "/api/v1/options", {
    params: {
      name: names.join(","),
      page: page,
      pageSize: pageSize
    }
  });
};

export const createOptions = (options: Option[]) => {
  return http.request<OptionsResult>("post", "/api/v1/options", {
    data: {
      options: options
    }
  });
};

export const updateOption = (option: Option) => {
  return http.request<OptionResult>("put", `/api/v1/options/${option.id}`, {
    data: {
      ...option
    }
  });
};

export const deleteOption = (option: Option) => {
  return http.request<OptionResult>("delete", `/api/v1/options/${option.id}`, {
    data: {
      name: option.name
    }
  });
};
