import { z } from 'zod'
import { ContactSchema } from './contact'
import { AddressSchema } from './address'

export const OfferShortSchema = z.object({
  id: z.uuid(),
  createdAt: z.coerce.date(),
  name: z.string().min(1),
})

export type OfferShort = z.infer<typeof OfferShortSchema>

export const OfferSchema = z.object({
  enterpriseName: z.string(),
  contact: ContactSchema.optional(),
  address: AddressSchema.optional(),
  techRequirements: z.array(z.string()).optional(),
  candidateScore: z.number().optional(),
  howToApply: z.string().optional(),
  description: z.string().optional(),
  agentRecap: z.string().optional(),
  name: z.string(),
})

export type Offer = z.infer<typeof OfferSchema>
