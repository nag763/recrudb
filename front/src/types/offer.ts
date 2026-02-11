import { z } from 'zod';
import { ContactSchema } from './contact';
import { AddressSchema } from './address';

export const OfferShortSchema = z.object({
  id: z.uuid(),
  createdAt: z.coerce.date(),
  name: z.string().optional(),
  enterpriseName: z.string().optional(),
  address: AddressSchema.optional(),
  candidateScore: z.number().optional(),
});

export type OfferShort = z.infer<typeof OfferShortSchema>;

export const OfferSchema = z.object({
  enterpriseName: z.string().optional(),
  contact: ContactSchema.optional(),
  address: AddressSchema.optional(),
  techRequirements: z.array(z.string()).optional(),
  candidateScore: z.number().optional(),
  howToApply: z.string().optional(),
  description: z.string().optional(),
  agentRecap: z.string().optional(),
  name: z.string().optional(),
  workLocation: z.string().optional(),
  latestNews: z.array(z.string()).optional(),
});

export type Offer = z.infer<typeof OfferSchema>;
