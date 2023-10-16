import { z } from "zod";

export const postSchema = z.object({
  topic: z.string(),
  description: z.string(),
  period: z.string(),
  howTo: z.string(),
  openPositions: z.array(z.string()),
  requiredSkills: z.array(z.string()),
  benefits: z.array(z.string()),
});

export type Post = z.infer<typeof postSchema>;
