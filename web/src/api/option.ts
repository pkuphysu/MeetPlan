import axios from "./axios";

export const getTermDateRange = () => {
  return axios.get<GetTermDateRangeResponse>('/api/v1/termdate')
}

export const getOption = (req: GetOptionRequest) => {
  return axios.get<GetOptionResponse>('/api/v1/option', req)
}

