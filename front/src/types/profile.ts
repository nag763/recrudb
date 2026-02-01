import { z } from "zod";
import { ContactSchema } from "./contact";
import { AddressSchema } from "./address";

const ExperienceSchema = z.object({
    startDate: z.coerce.date().optional(),
    endDate: z.coerce.date().optional(),
    company: z.string().optional(),
    mission: z.string().optional(), // Maps to your json:"mission" tag
    role: z.string().optional(),
    skills: z.array(z.string()).optional(),
});

const LanguageSchema = z.object({
    name: z.string().optional(),
    level: z.string().optional(),
});

const SocialSchema = z.object({
    name: z.string().optional(),
    link: z.url().optional(),
});

export const ProfileSchema = z.object({
    firstName: z.string(),
    lastName: z.string().optional(),
    contact: ContactSchema.optional(),
    skills: z.array(z.string()).optional(),
    lookingFor: z.string().optional(),
    experiences: z.array(ExperienceSchema).optional(),
    languages: z.array(LanguageSchema).optional(),
    socials: z.array(SocialSchema).optional(),
    location: AddressSchema.optional(),
});

export type Profile = z.infer<typeof ProfileSchema>;