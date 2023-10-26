"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getReviewResponseSchema } from "@/types/review-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function getReview(
  id: string,
  accessToken?: string
): Promise<z.infer<typeof getReviewResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(
    `${env.API_URL}/v1/reviews/${id}?accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getReviewResponseSchema, await response.json());
}
