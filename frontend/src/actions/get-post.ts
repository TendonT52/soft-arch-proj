"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getPostResponseSchema } from "@/types/post-service";
import { parseType } from "@/lib/utils";

export async function getPost(
  id: string,
  accessToken: string
): Promise<z.infer<typeof getPostResponseSchema>> {
  const response = await fetch(
    `${env.API_URL}/v1/posts/${id}?accessToken=${accessToken}`
  );

  return parseType(getPostResponseSchema, await response.json());
}
