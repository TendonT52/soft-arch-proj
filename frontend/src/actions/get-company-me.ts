"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getCompanyMeResponseSchema } from "@/types/user-service";
import { parseType } from "@/lib/utils";

export async function getCompanyMe(
  accessToken: string
): Promise<z.infer<typeof getCompanyMeResponseSchema>> {
  const response = await fetch(
    `${env.API_URL}/v1/company-me?accessToken=${accessToken}`
  );

  return parseType(getCompanyMeResponseSchema, await response.json());
}
