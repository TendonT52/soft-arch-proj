"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getPostsMeResponseSchema } from "@/types/post-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function getPostsMe(
  accessToken?: string
): Promise<z.infer<typeof getPostsMeResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(
    `${env.API_URL}/v1/posts/me?accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getPostsMeResponseSchema, await response.json());
}
