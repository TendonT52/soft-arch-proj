import { z } from "zod";

export const studentSchema = z.object({
  id: z.string(),
  name: z.string(),
  email: z.string(),
  description: z.string(),
  faculty: z.string(),
  major: z.string(),
  year: z.number(),
});

export type Student = z.infer<typeof studentSchema>;
