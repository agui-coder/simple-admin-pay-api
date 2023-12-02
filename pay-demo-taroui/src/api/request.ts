import Taro from '@tarojs/taro';

export interface RequestOptions {
  url: string;
  method?: 'GET' | 'POST' | 'PUT' | 'DELETE';
  data?: Record<string, any>;
  params?: Record<string, any>;
}

const baseUrl = "http://127.0.0.1:9107";

const defHttp = {
  async request(options: RequestOptions) {
    const { url, method = 'GET', data, params } = options;

    try {
      const response = await Taro.request({
        url: baseUrl + url,
        method,
        data,
        // Instead of using `params` directly, pass it as a query string in the URL
        // You might need to adjust this depending on Taro's request implementation
        header: {
          'content-type': 'application/json',
        },
        ...(method === 'GET' && params ? { data: params } : {}),
      });

      // Handle response status, you can customize this based on your needs
      if (response.statusCode >= 200 && response.statusCode < 300) {
        return response.data;
      } else {
        // Handle error cases, throw an error or return an object with error information
        throw new Error(`Request failed with status ${response.statusCode}`);
      }
    } catch (error) {
      // Handle network errors or other exceptions
      throw new Error(`Request failed: ${error.message}`);
    }
  },

  get(options: RequestOptions) {
    return this.request({ ...options, method: 'GET' });
  },

  post(options: RequestOptions) {
    return this.request({ ...options, method: 'POST' });
  },

  put(options: RequestOptions) {
    return this.request({ ...options, method: 'PUT' });
  },

  delete(options: RequestOptions) {
    return this.request({ ...options, method: 'DELETE' });
  },
};

export default defHttp;
