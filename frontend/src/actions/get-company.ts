import { env } from "env.mjs";
import { type z } from "zod";
import { getCompanyResponse } from "@/types/user-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export async function getCompany(
  id: string,
  accessToken?: string
): Promise<z.infer<typeof getCompanyResponse>> {
  const session = await getServerSession();
  if (!session) {
    throw new Error("No Session");
  }

  const response = await fetch(
    `${env.API_URL}/v1/company/${id}?accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getCompanyResponse, await response.json());
}
