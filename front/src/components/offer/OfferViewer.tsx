import { useQuery } from '@tanstack/react-query';
import { useContext, useMemo } from 'react';
import ReactMarkdown from 'react-markdown';
import { Skeleton } from '../ui/skeleton';
import OfferCards from './OfferCards';
import { getOffer } from '@/interactions/offers';
import { OfferContext } from '@/context/OfferContext';
import { cn } from '@/lib/utils';
import OfferExtended from './OfferExtended';

export default function OfferViewer() {
  const selectedOffer = useContext(OfferContext);

  const { data, isPending } = useQuery({
    queryFn: () => getOffer(selectedOffer?.id),
    queryKey: ['offer-id', selectedOffer?.id],
    enabled: !!selectedOffer?.id,
  });

  const offerSelected = useMemo(() => selectedOffer?.id != undefined, [selectedOffer]);
  const offer = useMemo(() => data?.data, [data]);

  return (
    <div className="flex justify-between w-full">
      {/* Offer grid */}
      <div className={cn("w-full max-h-screen overflow-scroll grid grid-cols-3 p-10 gap-4 text-center", selectedOffer && "p-6")}>
        <OfferCards />
      </div>
      {offerSelected && isPending && (
        <Skeleton className="max-h-screen w-full p-10 overflow-auto max-w-4xl mx-auto bg-white shadow-lg rounded-xl border border-gray-100" />
      )}
      {offerSelected && offer && <OfferExtended offer={offer}/>}
    </div>
  );
}
