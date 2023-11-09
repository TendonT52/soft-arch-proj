"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import {
  updateStudentResponseSchema,
  type updateStudentSchema,
} from "@/types/user-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function updateStudent(
  body: z.infer<typeof updateStudentSchema>
): Promise<z.infer<typeof updateStudentResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(`${env.API_URL}/v1/student`, {
    method: "PUT",
    body: JSON.stringify({
      ...body,
      accessToken: body.accessToken ?? session.accessToken,
    }),
  });

  return parseType(updateStudentResponseSchema, await response.json());
}
