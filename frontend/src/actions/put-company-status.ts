"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import {
  updateCompanyStatusResponseSchema,
  type updateCompanyStatusSchema,
} from "@/types/user-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export default async function putCompanyStatus(
  body: z.infer<typeof updateCompanyStatusSchema>
): Promise<z.infer<typeof updateCompanyStatusResponseSchema>> {
  const session = await getServerSession();
  if (!session) {
    throw new Error("No Session!!");
  }
  const response = await fetch(`${env.API_URL}/v1/company/status`, {
    method: "PUT",
    body: JSON.stringify({
      ...body,
      accessToken: body.accessToken ?? session.accessToken,
    }),
  });

  return parseType(updateCompanyStatusResponseSchema, await response.json());
}
