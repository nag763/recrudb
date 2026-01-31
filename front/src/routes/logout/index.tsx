import { auth } from '@/lib/gcpAuth';
import { signOut } from '@firebase/auth';
import { createFileRoute, useNavigate } from '@tanstack/react-router'

export const Route = createFileRoute('/logout/')({
  component: RouteComponent,
})

function RouteComponent() {
  const navigate = useNavigate();

  signOut(auth).finally(() => {
    navigate({
      to: "/"
    }) }
  );
}
