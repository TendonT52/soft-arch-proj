import decode from "jwt-decode";
import { z } from "zod";
import { UserRole } from "@/types/base/user";
import { parseType } from "./utils";

const accessTokenPayloadSchema = z.object({
  exp: z.number(),
  iat: z.number(),
  nbf: z.number(),
  role: z.nativeEnum(UserRole),
  userId: z.number(),
});

export function validateAccessToken(token: string) {
  const decoded = decode(token);
  const accessToken = parseType(accessTokenPayloadSchema, decoded);
  return Date.now() + 15 * 60 * 60 < accessToken.exp * 1000;
}

export function verifyAccessToken(token: string) {
  const decoded = decode(token);
  return parseType(accessTokenPayloadSchema, decoded);
}

export function getRefreshToken(response: Response) {
  const regex = /refreshToken=([^;]*)/;
  const setCookie = response.headers.getSetCookie();
  const match = setCookie[0]?.match(regex);
  if (match && match[1]) {
    return match[1];
  }
  return null;
}
