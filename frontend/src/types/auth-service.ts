/* AuthService */

import { z } from "zod";

/**
 * `POST /v1/admin`
 */
export const createAdminSchema = z.object({
  email: z.string(),
  password: z.string(),
  passwordConfirm: z.string(),
  accessToken: z.string(),
});

export const createAdminResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  id: z.string().optional(),
});

/**
 * `POST /v1/company`
 */
export const createCompanySchema = z.object({
  name: z.string(),
  email: z.string(),
  password: z.string(),
  passwordConfirm: z.string(),
  description: z.string(),
  location: z.string(),
  phone: z.string(),
  category: z.string(),
});

export const createCompanyResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  id: z.string(),
});

/**
 * `POST /v1/login`
 */
export const loginSchema = z.object({
  email: z.string(),
  password: z.string(),
});

export const loginResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  accessToken: z.string(),
  refreshToken: z.string(),
});

/**
 * `POST /v1/logout`
 */
export const logoutSchema = z.object({
  refreshToken: z.string(),
});

export const logoutResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
});

/**
 * `POST /v1/refresh`
 */
export const refreshSchema = z.object({
  refreshToken: z.string(),
});

export const refreshResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  accessToken: z.string(),
});

/**
 * `POST /v1/student`
 */
export const createStudentSchema = z.object({
  name: z.string(),
  email: z.string(),
  password: z.string(),
  passwordConfirm: z.string(),
  description: z.string(),
  faculty: z.string(),
  major: z.string(),
  year: z.number(),
});

export const createStudentResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
  id: z.string().optional(),
});

/**
 * `POST /v1/verify`
 */
export const verifySchema = z.object({
  code: z.string(),
  studentId: z.string(),
});

export const verifyResponseSchema = z.object({
  status: z.string(),
  message: z.string(),
});
