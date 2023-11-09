import { z } from "zod";

export const reviewSchema = z.object({
  title: z.string(),
  description: z.string(),
  rating: z.number(),
});

export type Review = z.infer<typeof reviewSchema>;
