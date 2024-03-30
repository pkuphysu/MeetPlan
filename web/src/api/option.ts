import { http } from "@/utils/http";

export type Option = {
  id?: string;
  name?: string;
  value?: string;
};

export type OptionResult = {
  data: Option;
};

export type OptionsResult = {
  data?: Option[];
};

export const searchOptions = (search?: string) => {
  return http.request<OptionsResult>("get", "/api/v1/options", {
    params: {
      search: search
    }
  });
};

export const getOption = (names: string[]) => {
  return http.request<OptionsResult>("get", "/api/v1/options", {
    params: {
      name: names.join(",")
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
