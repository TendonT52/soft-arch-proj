import { env } from "env.mjs";
import { type z } from "zod";
import {
  createReportResponseSchema,
  type createReportSchema,
} from "@/types/report-service";
import { getServerSession } from "@/lib/auth";
import { parseType } from "@/lib/utils";

export default async function createReport(
  body: z.infer<typeof createReportSchema>
): Promise<z.infer<typeof createReportResponseSchema>> {
  const session = await getServerSession();
  if (!session) {
    throw new Error("No Session!!");
  }

  const response = await fetch(`${env.API_URL}/v1/reports`, {
    method: "POST",
    body: JSON.stringify({
      ...body,
      accessToken: body.accessToken ?? session.accessToken,
    }),
  });
  return parseType(createReportResponseSchema, await response.json());
}
