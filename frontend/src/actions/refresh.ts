"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import {
  refreshResponseSchema,
  type refreshSchema,
} from "@/types/auth-service";
import { parseType } from "@/lib/utils";

export async function refresh(
  formData: z.infer<typeof refreshSchema>
): Promise<z.infer<typeof refreshResponseSchema>> {
  const response = await fetch(`${env.API_URL}/v1/refresh`, {
    method: "POST",
    body: JSON.stringify(formData),
  });

  return parseType(refreshResponseSchema, await response.json());
}
