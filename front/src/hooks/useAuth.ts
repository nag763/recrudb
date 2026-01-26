import { useQuery } from '@tanstack/react-query';
import {  onAuthStateChanged } from '@firebase/auth'; // Using the modular SDK
import { auth } from "../lib/gcpAuth";

export function useUser() {

  return useQuery({
    queryKey: ['user'],
    queryFn: () => {
      return new Promise<any | null>((resolve) => {
        const unsubscribe = onAuthStateChanged(auth, (user) => {
          unsubscribe();
          resolve(user);
        });
      });
    },
    // Keep the user data fresh so we don't re-run the observer constantly
    staleTime: Infinity, 
  });
}