"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getReviewsResponseSchema } from "@/types/review-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function getReviews(
  accessToken?: string
): Promise<z.infer<typeof getReviewsResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(
    `${env.API_URL}/v1/reviews?accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getReviewsResponseSchema, await response.json());
}
