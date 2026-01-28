import { auth } from "@/lib/gcpAuth";
import axios from "axios";

const injectBearer = async (config) => {
  const user = auth.currentUser;

  if (user) {
    const token = await user.getIdToken();
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
}

const createClient = (contentType?: string) => axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  headers: {
    'Content-Type': contentType
  }
});

const apiClient = createClient('application/json');

apiClient.interceptors.request.use(injectBearer);

apiClient.interceptors.response.use((response) => {
  if (response?.data?.success) {
    response.data = response.data.data
  }
});

export default apiClient;