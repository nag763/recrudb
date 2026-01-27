import { useQuery } from "@tanstack/react-query";
import { useUser } from "./useAuth";
import apiClient from "@/lib/client";


export function useProfile() {
  const { data: user } = useUser();

  return useQuery({
    queryKey: ['profile-me'],
    queryFn: async () => apiClient.get('/profile/me'),
    enabled: !!user, 
  });
}