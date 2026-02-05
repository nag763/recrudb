import OfferCards from '@/components/offer/OfferCards';
import OfferExtended from '@/components/offer/OfferExtended';
import { Skeleton } from '@/components/ui/skeleton';
import { getOffer } from '@/interactions/offers';
import { cn } from '@/lib/utils';
import { useQuery } from '@tanstack/react-query';
import { createFileRoute, useSearch } from '@tanstack/react-router'
import { useMemo } from 'react';

export const Route = createFileRoute('/offer/')({
  component: RouteComponent,
  validateSearch: (search: Record<string, unknown>) => {
    return {
      id: (search.id as string) || undefined,
    }
  },
})

function RouteComponent() {
  const { id } = useSearch({from: "/offer/"})

  const { data, isPending } = useQuery({
    queryFn: () => getOffer(id),
    queryKey: ['offer-id', id],
    enabled: !!id,
  });

  const offer = useMemo(() => data?.data, [data]);

  return (
    <div className="flex justify-between w-full">
      {/* Offer grid */}
      <div className={cn("w-full max-h-screen overflow-auto flex flex-col xl:grid xl:grid-cols-2 2xl:grid-cols-3 p-10 gap-4 text-center", id && "p-6")}>
        <OfferCards />
      </div>
      {id && isPending && (
        <Skeleton className="max-h-screen w-full p-10 overflow-auto max-w-4xl mx-auto bg-white shadow-lg rounded-xl border border-gray-100" />
      )}
      {id && offer && <OfferExtended offer={offer}/>}
    </div>
  );
}
