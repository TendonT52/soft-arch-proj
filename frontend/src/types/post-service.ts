/* PostService */

import { z } from "zod";
import { postSchema } from "./base/post";

/**
 * TODO
 * `GET /v1/benefits`
 */

/**
 * TODO
 * `GET /v1/open_positions`
 */

/**
 * `GET /v1/posts`
 */
export const getPostsResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  posts: z
    .array(
      postSchema.extend({
        owner: z.object({
          id: z.string(),
          name: z.string(),
        }),
        postId: z.string(),
        updatedAt: z.string(),
      })
    )
    .optional(),
  total: z.string().optional(),
});

/**
 * `POST /v1/posts`
 */
export const createPostSchema = z.object({
  post: postSchema,
  accessToken: z.string().optional(),
});

export const createPostResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  id: z.string().optional(),
});

/**
 * `GET /v1/posts/me`
 */
export const getPostsMeResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  posts: z
    .array(
      postSchema.extend({
        postId: z.string(),
        updatedAt: z.string(),
      })
    )
    .optional(),
});

/**
 * `GET /v1/posts/{id}`
 */
export const getPostResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  post: postSchema
    .extend({
      owner: z.object({
        id: z.string(),
        name: z.string(),
      }),
      updatedAt: z.string(),
    })
    .optional(),
});

/**
 * `DELETE /v1/posts/{id}`
 */
export const deletePostResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
});

/**
 * `PUT /v1/posts/{id}`
 */
export const updatePostSchema = z.object({
  post: z.object({
    topic: z.string(),
    description: z.string(),
    period: z.string(),
    howTo: z.string(),
    openPositions: z.array(
      z.object({
        action: z.literal("SAME").or(z.literal("ADD")).or(z.literal("REMOVE")),
        value: z.string().optional(),
      })
    ),
    requiredSkills: z.array(
      z.object({
        action: z.literal("SAME").or(z.literal("ADD")).or(z.literal("REMOVE")),
        value: z.string().optional(),
      })
    ),
    benefits: z.array(
      z.object({
        action: z.literal("SAME").or(z.literal("ADD")).or(z.literal("REMOVE")),
        value: z.string().optional(),
      })
    ),
  }),
  accessToken: z.string().optional(),
});

export const updatePostResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
});

/**
 * TODO
 * `GET /v1/required_skills`
 */
