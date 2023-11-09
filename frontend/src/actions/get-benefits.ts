"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getBenefitsResponseSchema } from "@/types/post-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function getBenefits(
  search: string,
  accessToken?: string
): Promise<z.infer<typeof getBenefitsResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(
    `${env.API_URL}/v1/benefits?search=${search}&accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getBenefitsResponseSchema, await response.json());
}
