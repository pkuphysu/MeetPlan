import axios, {AxiosError, AxiosInstance, AxiosRequestConfig, AxiosResponse, InternalAxiosRequestConfig} from 'axios';

interface ResultData<T> {
  code: number;
  message: string;
  data?: T;
}

class Request {
  private service: AxiosInstance;

  public constructor(config: AxiosRequestConfig) {
    this.service = axios.create(config);

    /**
     * 请求拦截
     */
    this.service.interceptors.request.use(
      (config: InternalAxiosRequestConfig) => {
        if (config!.url?.indexOf('login') === -1) {
          return config;
        }

        const token = localStorage.getItem("token") || '';
        if (token) {
          config!.headers!.Authorization = 'Bearer ' + token;
        }
        return config;
      },
      (error: AxiosError) => {
        return Promise.reject(error);
      }
    );

    /**
     * 响应拦截
     */
    this.service.interceptors.response.use(
      (response: AxiosResponse) => {
        const {data, status} = response;
        if (status === 401) {
          localStorage.removeItem("token");
          return Promise.reject('need login')
        }
        if (status !== 200 || data?.code !== 0) {
          return Promise.reject(data?.message || this.handleStatusCode(status));
        }
        return data.data;
      },
      (error: AxiosError) => {
        const {response} = error;
        if (response) {
          return Promise.reject(this.handleStatusCode(response.status));
        }
        if (!window.navigator.onLine) {
          return Promise.reject({code: 500, message: '网络连接失败'});
        }
      },
    );
  }

  handleStatusCode(code: number): string {
    switch (code) {
      case 403:
        return 'permission denied'
      case 404:
        return 'not found'
      case 500:
        return 'server error'
      default:
        return 'unknown error'
    }
  }

  get<T>(url: string, params?: any): Promise<T> {
    return this.service.get(url, {params});
  }

  post<T>(url: string, data?: any): Promise<T> {
    return this.service.post(url, data);
  }

  put<T>(url: string, data?: any): Promise<T> {
    return this.service.put(url, data);
  }

  delete<T>(url: string, data?: any): Promise<T> {
    return this.service.delete(url, data);
  }
}

export default new Request({
  baseURL: import.meta.env.BASE_URL,
  timeout: 3000,
})
