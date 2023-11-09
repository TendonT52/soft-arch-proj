"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getRequiredSkillsResponseSchema } from "@/types/post-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function getRequiredSkills(
  search: string,
  accessToken?: string
): Promise<z.infer<typeof getRequiredSkillsResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(
    `${env.API_URL}/v1/required_skills?search=${search}&accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getRequiredSkillsResponseSchema, await response.json());
}
