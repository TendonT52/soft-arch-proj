"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getReviewCompanyResponseSchema } from "@/types/review-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function getReviewCompany(
  cid: string,
  accessToken?: string
): Promise<z.infer<typeof getReviewCompanyResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(
    `${env.API_URL}/v1/reviews/company/${cid}?accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getReviewCompanyResponseSchema, await response.json());
}
