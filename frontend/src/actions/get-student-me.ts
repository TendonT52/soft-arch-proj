"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getStudentMeResponseSchema } from "@/types/user-service";
import { parseType } from "@/lib/utils";

export async function getStudentMe(
  accessToken: string
): Promise<z.infer<typeof getStudentMeResponseSchema>> {
  const response = await fetch(
    `${env.API_URL}/v1/student-me?accessToken=${accessToken}`
  );

  return parseType(getStudentMeResponseSchema, await response.json());
}
