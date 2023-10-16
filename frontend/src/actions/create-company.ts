"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import {
  createCompanyResponseSchema,
  type createCompanySchema,
} from "@/types/auth-service";
import { parseType } from "@/lib/utils";

export async function createCompany(
  formData: z.infer<typeof createCompanySchema>
): Promise<z.infer<typeof createCompanyResponseSchema>> {
  const response = await fetch(`${env.API_URL}/v1/company`, {
    method: "POST",
    body: JSON.stringify(formData),
  });

  // TODO: revalidate something
  return parseType(createCompanyResponseSchema, await response.json());
}
