"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getOpenPositionsResponseSchema } from "@/types/post-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function getOpenPositions(
  search: string,
  accessToken?: string
): Promise<z.infer<typeof getOpenPositionsResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(
    `${env.API_URL}/v1/open_positions?search=${search}&accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getOpenPositionsResponseSchema, await response.json());
}
