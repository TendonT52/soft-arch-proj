import { z } from "zod";

export const reportSchema = z.object({
  topic: z.string(),
  type: z.string(),
  description: z.string(),
});

export type report = z.infer<typeof reportSchema>;
