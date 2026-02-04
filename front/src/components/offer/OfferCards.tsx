import { useQuery } from '@tanstack/react-query';
import { useContext } from 'react';
import OfferCard from './OfferCard';
import { getOffers } from '@/interactions/offers';
import { OfferContext } from '@/context/OfferContext';

export default function OfferCards() {
  const { data: offers, isPending: isOffersPending } = useQuery({
    queryKey: ['profile-offers'],
    queryFn: getOffers,
  });

  const selectedOffer = useContext(OfferContext);

  if (isOffersPending) {
    return <p>Loading...</p>;
  }
  return offers?.data.map((e) => (
    <OfferCard offer={e} selected={e.id == selectedOffer?.id} key={e.id} />
  ));
}
