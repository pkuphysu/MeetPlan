import axios from './axios';

export const login = (req: LoginRequest) => {
  return axios.post<LoginResponse>('/api/v1/login', req);
}

export const getSelf = () => {
  return axios.get<GetSelfResponse>('/api/v1/user/self');
}

export const getUser = (id: number) => {
  return axios.get<GetUserResponse>(`/api/v1/user/${id}`);
}

export const listUser = (params: ListUserRequest) => {
  return axios.get<ListUserResponse>('/api/v1/user', params);
}

export const createUser = (params: CreateUserRequest) => {
  return axios.post<CreateUserResponse>('/api/v1/user', params);
}

export const updateUser = (params: UpdateUserRequest) => {
  return axios.put<UpdateUserResponse>(`/api/v1/user/${params.id}`, params);
}
