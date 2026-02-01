import { auth } from "@/lib/gcpAuth";
import axios, { type AxiosResponse, type InternalAxiosRequestConfig } from "axios";
import type z from "zod";

declare module "axios" {
  export interface AxiosRequestConfig {
    schema?: z.ZodSchema;
  }
}

const injectBearer = async (config: InternalAxiosRequestConfig) => {
  const user = auth.currentUser;

  if (user) {
    const token = await user.getIdToken();
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
}

const parseResponse = (response: AxiosResponse): AxiosResponse<any, any, {}> => {
  if (response?.data?.success) {
    let data = response.data.data;
    const schema = response.config.schema;
    if (!schema) {
      return { ...response, data: response.data?.data };
    }
    const parsed = schema.safeParse(data)
    if (!parsed.success) {
      console.error({message: "error while parsing for schema", schema,data, error: parsed.error})
      throw parsed.error;
    }
    return { ...response, data: parsed?.data };
  }
  return response
};

const createClient = (contentType?: string) => axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  headers: {
    'Content-Type': contentType
  }
});

const apiClient = createClient('application/json');

apiClient.interceptors.request.use(injectBearer);

apiClient.interceptors.response.use(parseResponse);

export default apiClient;