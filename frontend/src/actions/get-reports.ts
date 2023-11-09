import { env } from "env.mjs";
import { type z } from "zod";
import { getReportsResponseSchema } from "@/types/report-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export default async function getReports(
  accessToken?: string
): Promise<z.infer<typeof getReportsResponseSchema>> {
  const session = await getServerSession();
  if (!session) {
    throw new Error("No Session!!");
  }
  const response = await fetch(
    `${env.API_URL}/v1/reports?accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getReportsResponseSchema, await response.json());
}
