import { createFileRoute, useNavigate } from '@tanstack/react-router';
import { useReducer } from 'react';
import { useProfile } from '@/hooks/useProfile';
import { SidebarProvider } from '@/components/ui/sidebar';
import { AppSidebar } from '@/components/AppSidebar';
import {
  OfferContext,
  OfferContextDispatcher,
  offerReducer,
} from '@/context/OfferContext';
import OfferViewer from '@/components/offer/OfferViewer';

export const Route = createFileRoute('/')({
  component: App,
});

function App() {
  const { data, isLoading } = useProfile();

  const navigate = useNavigate();

  const [offer, dispatch] = useReducer(offerReducer, null);

  if (isLoading) {
    return <p>Loading...</p>;
  }

  if (data?.status == 204) {
    navigate({ to: '/profile' });
  }

  return (
    <OfferContext value={offer}>
      <OfferContextDispatcher value={dispatch}>

        <OfferViewer />
      </OfferContextDispatcher>
    </OfferContext>
  );
}
