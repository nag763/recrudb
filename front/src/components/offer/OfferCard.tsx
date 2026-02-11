import { MapPin, TrendingUp } from 'lucide-react';
import { Badge } from '../ui/badge';
import type { OfferShort } from '@/types/offer';
import { cn } from '@/lib/utils';
import { useMemo } from 'react';
import { TooltipContent, Tooltip, TooltipTrigger } from '../ui/tooltip';
import { Link, useSearch } from '@tanstack/react-router';
import OfferContextMenu from './OfferContextMenu';

interface OfferCardProps {
  offer: OfferShort;
}

export default function OfferCard({ offer }: OfferCardProps) {

  const { id } = useSearch({ from: "/offer/" })

  const selected = useMemo(() => id == offer.id, [id]);
  return (

    <OfferContextMenu offerId={offer.id}>
      <Link
        to={"/offer"}
        search={{ id: selected ? undefined : offer.id }}
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
                {id ?
                  <Tooltip>
                    <TooltipTrigger asChild>
                      <TrendingUp />
                    </TooltipTrigger>
                    <TooltipContent>
                      <p>This offer matches your profile</p>
                    </TooltipContent>
                  </Tooltip> :
                  <span>High match</span>
                }
              </Badge>
            )}
          </div>

        </div>
        <div className="flex justify-between text-xs text-gray-400 items-center ">
          <div className="flex text-gray-600 ">
            <MapPin className="w-4 h-4 mr-1" />
            <span>{offer.address?.city || offer.address?.country}</span>
          </div>
          <p>{offer.createdAt.toLocaleDateString()}</p>
        </div>
      </Link>
    </OfferContextMenu>
  );
}
