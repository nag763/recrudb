import { MapPin } from 'lucide-react';
import { Badge } from '../ui/badge';
import type { OfferShort } from '@/types/offer';
import { cn } from '@/lib/utils';
import useSelectOffer from '@/hooks/useSelectOffer';
import { useMemo } from 'react';

interface OfferCardProps {
  offer: OfferShort;
  selected: boolean
}

export default function OfferCard({ offer }: OfferCardProps) {
  const {selectedOffer, onOfferSelected: onClick} = useSelectOffer();
  const selected = useMemo(() => selectedOffer?.id == offer.id, [selectedOffer] );
  return (
    <div
      onClick={() => onClick(offer.id)}
      key={offer.id}
      className={cn("bg-white border border-gray-200 rounded-lg p-4 shadow-sm hover:shadow-md cursor-pointer",
        "transition-shadow duration-200 flex flex-col justify-between",
         selected && "border-4 border-indigo-200")}
    >
      <div>
        <div className="flex justify-between items-start">
          <div className="text-left">
            <p className="text-sm text-gray-500">{offer.enterpriseName}</p>
            <h3 className="text-lg font-semibold text-gray-800">
              {offer.name}
            </h3>
          </div>
          {18 <= (offer?.candidateScore || 0) && (
            <Badge
              variant="outline"
              className="bg-green-100 text-green-800 border-green-200 text-xs"
            >
              High match
            </Badge>
          )}
        </div>
        <div className="flex items-center text-sm text-gray-600 mt-2">
          <MapPin className="w-4 h-4 mr-1" />
          <span>{offer.address?.city || offer.address?.country}</span>
        </div>
      </div>
      <div className="text-right text-xs text-gray-400 mt-4">
        <p>{offer.createdAt.toLocaleDateString()}</p>
      </div>
    </div>
  );
}
