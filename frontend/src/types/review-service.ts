/* ReviewService */

import { z } from "zod";
import { reviewSchema } from "./base/review";

/**
 * `GET /v1/reviews`
 */
export const getReviewsResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  reviews: z
    .array(
      reviewSchema.extend({
        id: z.string(),
        updatedAt: z.string(),
        company: z.object({
          id: z.string(),
          name: z.string(),
        }),
      })
    )
    .optional(),
  total: z.string().optional(),
});

/**
 * `POST /v1/reviews`
 */
export const createReviewSchema = z.object({
  review: reviewSchema.extend({
    cid: z.string(),
    isAnonymous: z.boolean(),
  }),
  accessToken: z.string().optional(),
});

export const createReviewResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  id: z.string().optional(),
});

/**
 * TODO
 * `GET /v1/reviews/{cid}`
 */

/**
 * TODO
 * `GET /v1/reviews/{id}`
 */

/**
 * TODO
 * `DELETE /v1/reviews/{id}`
 */

/**
 * TODO
 * `PUT /v1/reviews/{id}`
 */
