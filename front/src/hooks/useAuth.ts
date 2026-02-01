import { useQuery } from '@tanstack/react-query';
import {  onAuthStateChanged, type User } from '@firebase/auth'; // Using the modular SDK
import { auth } from "../lib/gcpAuth";

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