import z from 'zod';

export const ContactSchema = z.object({
  email: z.string().optional(),
  name: z.string().optional(),
  number: z.string().optional(),
});
