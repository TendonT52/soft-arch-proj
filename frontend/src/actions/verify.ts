import { env } from "env.mjs";
import { type z } from "zod";
import { verifyResponseSchema, type verifySchema } from "@/types/auth-service";
import { parseType } from "@/lib/utils";

export async function verify(
  body: z.infer<typeof verifySchema>
): Promise<z.infer<typeof verifyResponseSchema>> {
  const response = await fetch(`${env.API_URL}/v1/verify`, {
    method: "POST",
    body: JSON.stringify(body),
  });

  return parseType(verifyResponseSchema, await response.json());
}
