import { useQuery } from '@tanstack/react-query';
import { onAuthStateChanged } from '@firebase/auth';
import { auth } from '../lib/gcpAuth';
import type { User } from '@firebase/auth'; // Using the modular SDK

export function useUser() {
  return useQuery({
    queryKey: ['user'],
    queryFn: () => {
      return new Promise<User | null>((resolve) => {
        const unsubscribe = onAuthStateChanged(auth, (user) => {
          unsubscribe();
          resolve(user);
        });
      });
    },
    staleTime: 1000 * 60 * 5,
  });
}
