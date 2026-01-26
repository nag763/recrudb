import { useQuery } from "@tanstack/react-query";
import { useUser } from "./useAuth";

export function useApiData() {
  const { data: user } = useUser();

  return useQuery({
    queryKey: ['protected-data'],
    queryFn: async () => {
      const token = await user?.getIdToken();
      
      const response = await fetch('https://your-gcp-backend.a.run.app/data', {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      return response.json();
    },
    enabled: !!user, 
  });
}