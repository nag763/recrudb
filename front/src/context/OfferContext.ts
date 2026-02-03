import { createContext } from 'react';
import type { ActionDispatch } from 'react';

export const OfferContext = createContext<OfferContextType>({});
export const OfferContextDispatcher =
  createContext<OfferContextDispatcher>(null);

type OfferContextType = {
  id?: string;
} | null;

type OfferContextAction =
  | { type: 'selected'; id: string }
  | { type: 'default' };

type OfferContextDispatcher = ActionDispatch<
  [action: OfferContextAction]
> | null;

export function offerReducer(
  _prevState: OfferContextType,
  action: OfferContextAction,
): OfferContextType {
  switch (action.type) {
    case 'selected': {
      return { id: action.id };
    }
    default:
      return {};
  }
}
