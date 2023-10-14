"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import {
  createStudentResponseSchema,
  type createStudentSchema,
} from "@/types/auth-service";
import { parseType } from "@/lib/utils";

export async function createStudent(
  formData: z.infer<typeof createStudentSchema>
): Promise<z.infer<typeof createStudentResponseSchema>> {
  console.log("formData", formData);
  const response = await fetch(`${env.API_URL}/v1/student`, {
    method: "POST",
    body: JSON.stringify(formData),
  });
  const data = parseType(createStudentResponseSchema, await response.json());
  console.log("response", data);

  // TODO: revalidate something
  return data;
}
