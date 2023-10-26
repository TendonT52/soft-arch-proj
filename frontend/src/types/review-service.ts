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
  total: z.number().optional(),
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
 * `GET /v1/reviews/{id}`
 */
export const getReviewResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  review: reviewSchema.extend({
    id: z.string(),
    updatedAt: z.string(),
    owner: z.object({
      id: z.string().optional(),
      name: z.string(),
    }),
    company: z.object({
      id: z.string(),
      name: z.string(),
    }),
  }),
});

/**
 * TODO
 * `DELETE /v1/reviews/{id}`
 */
export const deleteReviewResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
});

/**
 * TODO
 * `PUT /v1/reviews/{id}`
 */
export const updateReviewSchema = z.object({
  review: z.object({
    title: z.string(),
    description: z.string(),
    rating: z.number(),
    isAnonymous: z.boolean(),
  }),
  accessToken: z.string().optional(),
});

export const updateReviewResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
});
