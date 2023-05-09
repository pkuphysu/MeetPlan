import axios from "./axios";

export interface TermDateRange{
  start: Number
  end: Number
}

export const getTermDateRange = () => {
  return axios.get<TermDateRange>('/api/v1/termdate')
}

export const getOption = (key: string) => {
  return axios.get<string>('/api/v1/option', {key: key})
}

