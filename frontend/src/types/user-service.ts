/* UserService */

import { z } from "zod";
import { companySchema } from "./base/company";
import { studentSchema } from "./base/student";

/**
 * `GET /v1/companies`
 */
export const getCompaniesResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  companies: z.array(companySchema).optional(),
  total: z.string().optional(),
});

/**
 * `GET /v1/companies/approved`
 */
export const getCompaniesApprovedResponseSchema = getCompaniesResponseSchema;

/**
 * `PUT /v1/company`
 */
export const updateCompanySchema = z.object({
  accessToken: z.string().optional(),
  company: z.object({
    name: z.string(),
    description: z.string(),
    location: z.string(),
    phone: z.string(),
    category: z.string(),
  }),
});

export const updateCompanyResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
});

/**
 * `GET /v1/company-me`
 */
export const getCompanyMeResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  company: companySchema,
});

/**
 * `PUT /v1/company/status`
 */
export const updateCompanyStatusSchema = z.object({
  accessToken: z.string().optional(),
  id: z.string(),
  status: z.string(),
});

export const updateCompanyStatusResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
});

/**
 * `GET /v1/company/{id}`
 */
export const getCompanyResponse = z.object({
  status: z.string(),
  message: z.string(),
  company: companySchema.optional(),
});

/**
 * `PUT /v1/student`
 */
export const updateStudentSchema = z.object({
  accessToken: z.string().optional(),
  student: z.object({
    name: z.string(),
    description: z.string(),
    faculty: z.string(),
    major: z.string(),
    year: z.number(),
  }),
});

export const updateStudentResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
});

/**
 * `GET /v1/student-me`
 */
export const getStudentMeResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  student: studentSchema,
});

/**
 * `GET /v1/student/{id}`
 */
export const getStudentResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  student: studentSchema,
});
