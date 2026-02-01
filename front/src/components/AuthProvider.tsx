import { onAuthStateChanged } from '@firebase/auth'
import { useQueryClient } from '@tanstack/react-query'
import { useEffect } from 'react'
import { auth } from '@/lib/gcpAuth'

export default function AuthProvider({
  children,
}: {
  children: React.ReactNode
}) {
  const queryClient = useQueryClient()

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, (user) => {
      queryClient.setQueryData(['user'], user)
    })

    return () => unsubscribe()
  }, [queryClient])

  return <>{children}</>
}
