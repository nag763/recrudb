import apiClient from "@/lib/client";
import { OfferSchema, OfferShortSchema, type Offer, type OfferShort } from "@/types/offer";
import z from "zod";
import type { ApiResponse } from "./type";

export const getOffers = (): ApiResponse<Array<OfferShort>> => (apiClient.get('/profile/me/offers', { schema: z.array(OfferShortSchema) }));
export const getOffer = (offerId?: string): ApiResponse<Offer> => (apiClient.get(`/profile/me/offers/${offerId}`, {schema: OfferSchema}))
export const deleteOffer = (offerId: string) => (apiClient.delete(`/profile/me/offers/${offerId}`)) 