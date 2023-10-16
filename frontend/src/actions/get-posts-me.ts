"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getPostsMeResponseSchema } from "@/types/post-service";
import { parseType } from "@/lib/utils";

export async function getPostsMe(
  accessToken: string
): Promise<z.infer<typeof getPostsMeResponseSchema>> {
  const response = await fetch(
    `${env.API_URL}/v1/posts/me?accessToken=${accessToken}`
  );

  return parseType(getPostsMeResponseSchema, await response.json());
}
