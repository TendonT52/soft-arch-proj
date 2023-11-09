"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getStudentResponseSchema } from "@/types/user-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function getStudent(
  id: string,
  accessToken?: string
): Promise<z.infer<typeof getStudentResponseSchema>> {
  const session = await getServerSession();
  if (!session) {
    throw new Error("No Session");
  }

  const response = await fetch(
    `${env.API_URL}/v1/student/${id}?accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getStudentResponseSchema, await response.json());
}
