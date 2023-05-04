import axios from './axios';

export interface LoginParams {
  code: string;
}

export type LoginResult = string;

export const login = (params: LoginParams) => {
  return axios.post<LoginResult>('/api/v1/login', params);
}

  export interface User {
    id: number;
    pku_id: string;
    name: string;
    email: string;
    is_teacher: boolean;
    is_admin: boolean;
    is_active: boolean;
    gender?: number;
    avatar?: string;
    phone?: string;
    department?: string;
    major?: string;
    grade?: string;
    dorm?: string;
    office?: string;
    introduction?: string;
    email_change?: string;
  }

  export interface ListUserParams {
    page_no: number;
    page_size: number;
    is_active?: boolean;
    is_teacher?: boolean;
    is_admin?: boolean;
    ids?: number[];
  }

  export interface CreateUserParams {
    pku_id: string;
    name: string;
    email: string;
    is_teacher?: boolean;
    is_admin?: boolean;
    gender?: number;
    avatar?: string;
    phone?: string;
    department?: string;
    major?: string;
    grade?: string;
    dorm?: string;
    office?: string;
    introduction?: string;
  }

  export interface UpdateUserParams {
    pku_id?: string;
    name?: string;
    email?: string;
    is_teacher?: boolean;
    is_admin?: boolean;
    is_active?: boolean;
    gender?: number;
    avatar?: string;
    phone?: string;
    department?: string;
    major?: string;
    grade?: string;
    dorm?: string;
    office?: string;
    introduction?: string;
  }


export const getSelf = () => {
  return axios.get<User>('/api/v1/user/self');
}

export const getUser = (id: number) => {
  return axios.get<User>(`/api/v1/user/${id}`);
}

export const listUser = (params: ListUserParams) => {
  return axios.get<User[]>('/api/v1/user', params);
}

export const createUser = (params: CreateUserParams) => {
  return axios.post<User>('/api/v1/user', params);
}

export const updateUser = (id: number, params: UpdateUserParams) => {
  return axios.put<User>(`/api/v1/user/${id}`, params);
}
