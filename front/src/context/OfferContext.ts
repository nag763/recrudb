import { createContext } from 'react';

export const OfferContext = createContext(null);
export const OfferContextDispatcher = createContext(null);

export function offerReducer(offer, action) {
    switch(action.type) {
        case 'selected': {
            return action.id
        }
        default:
            break
    }
}