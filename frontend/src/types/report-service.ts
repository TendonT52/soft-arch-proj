import { z } from "zod";
import { reportSchema } from "./base/report";

/**
 * POST /v1/reports/
 */
export const createReportSchema = z.object({
  report: reportSchema,
  accessToken: z.string().optional(),
});
export const createReportResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  id: z.string().optional(),
});
/**
 * GET /v1/reports
 */
export const getReportsResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  reports: z
    .array(
      reportSchema.extend({
        updatedAt: z.string(),
      })
    )
    .optional(),
  total: z.string(),
});
/**
 * GET /v1/reports/{id}
 */
export const getReportResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  report: reportSchema.extend({
    updatedAt: z.string(),
  }),
});
