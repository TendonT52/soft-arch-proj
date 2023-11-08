import { env } from "env.mjs";
import { type z } from "zod";
import { getReportResponseSchema } from "@/types/report-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export default async function getReports(
  id: string,
  accessToken?: string
): Promise<z.infer<typeof getReportResponseSchema>> {
  const session = await getServerSession();
  if (!session) {
    throw new Error("No Session!!");
  }
  const response = await fetch(
    `${env.API_URL}/v1/reports/${id}?accessToken=${
      accessToken ?? session.accessToken
    }`
  );

  return parseType(getReportResponseSchema, await response.json());
}
