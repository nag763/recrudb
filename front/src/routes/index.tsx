import { createFileRoute, useNavigate } from '@tanstack/react-router'
import { useProfile } from '@/hooks/useProfile';
import { SidebarProvider } from '@/components/ui/sidebar';
import { ChatSidebar } from '@/components/ChatsSidebar';
import { useReducer } from 'react';
import { OfferContext, OfferContextDispatcher, offerReducer } from '@/context/OfferContext';
import OfferViewer from '@/components/OfferViewer';

export const Route = createFileRoute('/')({
  component: App,
})

function App() {
  const { data, isLoading } = useProfile();

  const navigate = useNavigate()

  const [offer, dispatch] = useReducer(offerReducer, null);

  if (isLoading) {
    return <p>Loading...</p>
  }

  if (data?.status == 204) {
    navigate({ to: '/profile' })
  }

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
