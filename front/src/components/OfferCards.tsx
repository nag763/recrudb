import { useQuery } from '@tanstack/react-query';
import { useCallback, useContext } from 'react';
import OfferCard from './OfferCard';
import { getOffers } from '@/interactions/offers';
import { OfferContext, OfferContextDispatcher } from '@/context/OfferContext';

export default function OfferCards() {
  const { data: offers, isPending: isOffersPending } = useQuery({
    queryKey: ['profile-offers'],
    queryFn: getOffers,
  });

  const selectedOffer = useContext(OfferContext);
  const dispatch = useContext(OfferContextDispatcher);

  const itemClicked = useCallback(
    (id: string) => {
      if (selectedOffer?.id == id) {
        dispatch && dispatch(
          {
            type: "default"
          }
        )
        return;
      }
      dispatch &&
        dispatch({
          type: 'selected',
          id: id,
        });
    },
    [dispatch, selectedOffer],
  );

  if (isOffersPending) {
    return <p>Loading...</p>;
  }
  return offers?.data.map((e) => (
    <OfferCard offer={e} onClick={itemClicked} key={e.id} />
  ));
}
