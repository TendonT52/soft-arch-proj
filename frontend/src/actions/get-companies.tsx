import { env } from "env.mjs";
import { type z } from "zod";
import { getCompaniesResponseSchema } from "@/types/user-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export default async function getCompanies(
  accessToken?: string
): Promise<z.infer<typeof getCompaniesResponseSchema>> {
  const session = await getServerSession();
  if (!session) {
    throw new Error("No Session!!");
  }
  const response = await fetch(
    `${env.API_URL}/v1/companies?accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getCompaniesResponseSchema, await response.json());
}
