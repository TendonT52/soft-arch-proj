"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import {
  createStudentResponseSchema,
  type createStudentSchema,
} from "@/types/auth-service";
import { parseType } from "@/lib/utils";

export async function createStudent(
  body: z.infer<typeof createStudentSchema>
): Promise<z.infer<typeof createStudentResponseSchema>> {
  const response = await fetch(`${env.API_URL}/v1/student`, {
    method: "POST",
    body: JSON.stringify(body),
  });

  // TODO: revalidate something
  return parseType(createStudentResponseSchema, await response.json());
}
