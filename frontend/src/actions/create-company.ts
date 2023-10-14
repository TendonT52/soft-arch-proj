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
  console.log("formData", formData);
  const response = await fetch(`${env.API_URL}/v1/company`, {
    method: "POST",
    body: JSON.stringify(formData),
  });
  const data = parseType(createCompanyResponseSchema, await response.json());
  console.log("response", data);

  // TODO: revalidate something
  return data;
}
