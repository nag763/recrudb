import apiClient from "@/lib/client";

export const getOffers = () => (apiClient.get('/profile/me/offers'));
export const deleteOffer = (offerId: string) => (apiClient.delete(`/profile/me/offers/${offerId}`)) 