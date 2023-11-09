import { z } from "zod";

export const companySchema = z.object({
  id: z.string(),
  name: z.string(),
  email: z.string(),
  description: z.string(),
  location: z.string(),
  phone: z.string(),
  category: z.string(),
  status: z.string(),
});

export type Company = z.infer<typeof companySchema>;
