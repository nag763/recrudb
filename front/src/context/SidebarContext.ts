import { createContext } from 'react';
import type { ActionDispatch } from 'react';

type SidebarOptions = 'offers' | 'chats'

export const OfferContext = createContext<SidebarOptions>('offers');
export const OfferContextDispatcher =
    createContext<SidebarContextDispatcher>(null);

type SidebarContextAction =
    | { type: 'select'; option: SidebarOptions }
    | { type: 'default' };

type SidebarContextDispatcher = ActionDispatch<
    [action: SidebarContextAction]
> | null;

export function offerReducer(
    _prevState: SidebarOptions,
    action: SidebarContextAction,
): SidebarOptions {
    switch (action.type) {
        case 'select': {
            return action.option;
        }
        default:
            return 'offers';
    }
}
