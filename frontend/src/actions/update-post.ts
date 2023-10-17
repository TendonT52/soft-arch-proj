"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import {
  updatePostResponseSchema,
  type updatePostSchema,
} from "@/types/post-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function updatePost(
  id: string,
  body: z.infer<typeof updatePostSchema>
): Promise<z.infer<typeof updatePostResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(`${env.API_URL}/v1/posts/${id}`, {
    method: "PUT",
    body: JSON.stringify({
      ...body,
      accessToken: session.accessToken,
    }),
  });

  return parseType(updatePostResponseSchema, await response.json());
}
