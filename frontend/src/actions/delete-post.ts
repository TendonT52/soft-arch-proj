"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { deletePostResponseSchema } from "@/types/post-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function deletePost(
  id: string,
  accessToken?: string
): Promise<z.infer<typeof deletePostResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(
    `${env.API_URL}/v1/posts/${id}?accessToken=${
      accessToken ?? session.accessToken
    }`,
    {
      method: "DELETE",
    }
  );

  return parseType(deletePostResponseSchema, await response.json());
}
