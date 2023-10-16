"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { loginResponseSchema, type loginSchema } from "@/types/auth-service";
import { getRefreshToken } from "@/lib/token";
import { parseType } from "@/lib/utils";

export async function login(
  formData: z.infer<typeof loginSchema>
): Promise<z.infer<typeof loginResponseSchema>> {
  const response = await fetch(`${env.API_URL}/v1/login`, {
    method: "POST",
    body: JSON.stringify(formData),
  });
  const data = (await response.json()) as object;
  const refreshToken = getRefreshToken(response);

  // TODO: revalidate something
  return parseType(loginResponseSchema, { ...data, refreshToken });
}
