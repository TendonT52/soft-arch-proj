"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import {
  createReviewResponseSchema,
  type createReviewSchema,
} from "@/types/review-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function createReview(
  body: z.infer<typeof createReviewSchema>
): Promise<z.infer<typeof createReviewResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(`${env.API_URL}/v1/reviews`, {
    method: "POST",
    body: JSON.stringify({
      ...body,
      accessToken: body.accessToken ?? session.accessToken,
    }),
  });

  return parseType(createReviewResponseSchema, await response.json());
}
