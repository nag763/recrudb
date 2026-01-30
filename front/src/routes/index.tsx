import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { useUser } from '@/hooks/useAuth'
import { useProfile } from '@/hooks/useProfile';
import { SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar';
import { ChatSidebar } from '@/components/ChatsSidebar';
import { useReducer } from 'react';
import { OfferContext, OfferContextDispatcher, offerReducer } from '@/context/OfferContext';
import OfferViewer from '@/components/OfferViewer';

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

  const [offer, dispatch] = useReducer(offerReducer, null);

  return (
    <OfferContext value={offer}>
      <OfferContextDispatcher value={dispatch}>
        <SidebarProvider>
          <ChatSidebar/>

          <OfferViewer />
        </SidebarProvider >
      </OfferContextDispatcher>
    </OfferContext>
  )
}
