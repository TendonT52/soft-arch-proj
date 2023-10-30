"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import {
  updateCompanyResponseSchema,
  type updateCompanySchema,
} from "@/types/user-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function updateCompany(
  body: z.infer<typeof updateCompanySchema>
): Promise<z.infer<typeof updateCompanyResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(`${env.API_URL}/v1/company`, {
    method: "PUT",
    body: JSON.stringify({
      ...body,
      accessToken: body.accessToken ?? session.accessToken,
    }),
  });

  return parseType(updateCompanyResponseSchema, await response.json());
}
