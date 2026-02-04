import { OfferContext, OfferContextDispatcher } from "@/context/OfferContext";
import { useCallback, useContext } from "react";

function useSelectOffer() {
    const selectedOffer = useContext(OfferContext);
    const dispatch = useContext(OfferContextDispatcher);

    return {selectedOffer, onOfferSelected: useCallback((id: string) => {
        if (!dispatch) return;

        if (selectedOffer?.id === id) {
            dispatch({ type: "default" });
        } else {
            dispatch({ type: 'selected', id });
        }
    }, [dispatch, selectedOffer])};
}

export default useSelectOffer;