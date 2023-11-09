"use server";

import { env } from "env.mjs";
import { type z } from "zod";
import { getPostsResponseSchema } from "@/types/post-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function getPosts(
  accessToken?: string,
  searchOptions?: {
    searchCompany?: string;
    searchOpenPosition?: string;
    searchRequiredSkill?: string;
    searchBenefit?: string;
  }
): Promise<z.infer<typeof getPostsResponseSchema>> {
  const session = await getServerSession(); // this will retrieve new access token if it's expired
  if (!session) {
    throw new Error("No session");
  }

  const response = await fetch(
    `${env.API_URL}/v1/posts?accessToken=${
      accessToken ?? session.accessToken
    }&searchOptions.searchCompany=${searchOptions?.searchCompany}&searchOptions.searchOpenPosition=${searchOptions?.searchOpenPosition}&searchOptions.searchRequiredSkill=${searchOptions?.searchRequiredSkill}&searchOptions.searchBenefit=${searchOptions?.searchBenefit}`
  );

  return parseType(getPostsResponseSchema, await response.json());
}
