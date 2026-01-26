import { createFileRoute } from '@tanstack/react-router'
import { LoginComponent } from '@/components/LoginComponent'
import { useUser } from '@/hooks/useAuth'
import { auth } from '@/lib/gcpAuth';
import { useApiData } from '@/hooks/useApiData';

export const Route = createFileRoute('/')({
  component: App,
})

function App() {
  const { data: user, isLoading } = useUser();
  const { data } = useApiData();

  if (isLoading) return <div>Checking session...</div>;
  return (
    <div className="text-center">
      {user ? (
        <div>
          <span>Welcome, {user.displayName}</span>
          <button onClick={() => auth.signOut()}>Logout</button>
        </div>
      ) : (
        <LoginComponent />
      )}
    </div>
  )
}
