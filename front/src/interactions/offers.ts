import apiClient from "@/lib/client";

export const getOffers = () => (apiClient.get('/profile/me/offers'));