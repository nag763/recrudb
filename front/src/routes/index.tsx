import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { useUser } from '@/hooks/useAuth'
import { useProfile } from '@/hooks/useProfile';
import { SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar';
import { ChatSidebar } from '@/components/ChatsSidebar';

export const Route = createFileRoute('/')({
  component: App,
})

function App() {
  const { data: user, isLoading } = useUser();
  const { data } = useProfile();

  const navigate = useNavigate()

  if (isLoading) {
    return <p>Loading...</p>
  }

  if (data?.status == 204) {
    navigate({ to: '/profile' })
  }

  return (
    <SidebarProvider>
      <ChatSidebar>
        <div className="text-center">
          <div>
          </div>
        </div>

      </ChatSidebar>
    </SidebarProvider >
  )
}
