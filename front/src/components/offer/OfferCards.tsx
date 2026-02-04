import { useQuery } from '@tanstack/react-query';
import OfferCard from './OfferCard';
import { getOffers } from '@/interactions/offers';
import { useMemo } from 'react';

export default function OfferCards() {
  const { data, isPending: isOffersPending } = useQuery({
    queryKey: ['profile-offers'],
    queryFn: getOffers,
  });
  const offers = useMemo(() => data?.data, [data]);


  if (isOffersPending) {
    return <p>Loading...</p>;
  }


  if(!offers) {
    return <div className='w-full h-full flex flex-col gap-12 items-center align-middle justify-center col-span-full'>
      <img src="/decide.svg" height={400} width={400}/>
      <h1 className='text-4xl'>No offer so far !</h1>
      <h2>Use the button on the left panel to add a new one.</h2>
    </div>
  }

  return offers?.map((e) => (
    <OfferCard offer={e} key={e.id} />
  ));
}
