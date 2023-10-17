import { cx } from "cva";
import { type ClassValue } from "cva/types";
import { twMerge } from "tailwind-merge";
import { type z } from "zod";

export function cn(...inputs: ClassValue[]) {
  return twMerge(cx(inputs));
}

export function formatDate(input: string | number): string {
  const date = new Date(input);
  return date.toLocaleDateString("en-US", {
    month: "long",
    day: "numeric",
    year: "numeric",
  });
}

export async function sleep(ms: number) {
  await new Promise((resolve) => {
    setTimeout(resolve, ms);
  });
}

export function parseType<T>(schema: z.ZodSchema<T>, data: unknown) {
  return schema.parse(data);
}

export function getSearchArray(searchParam?: string | string[]) {
  return searchParam
    ? Array.isArray(searchParam)
      ? searchParam
      : [searchParam]
    : [];
}
