import { cx } from "cva";
import { type ClassValue } from "cva/types";
import { format } from "date-fns";
import { type DateRange } from "react-day-picker";
import { twMerge } from "tailwind-merge";
import { type z } from "zod";

export function cn(...inputs: ClassValue[]) {
  return twMerge(cx(inputs));
}

export function formatDate(input: string | number) {
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

export function formatPeriod(date?: DateRange) {
  return date?.from
    ? date.to
      ? `${format(date.from, "LLL dd, y")} - ${format(date.to, "LLL dd, y")}`
      : format(date.from, "LLL dd, y")
    : "Date range";
}

export function parsePeriod(period: string) {
  const regex =
    /(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)\s\d{1,2},\s\d{4}\s-\s(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)\s\d{1,2},\s\d{4}/g;
  const match = period.match(regex);
  if (!match || !match[0]) return undefined;
  const [from, to] = match[0].split(" - ");
  return {
    from: from ? new Date(from) : undefined,
    to: to ? new Date(to) : undefined,
  };
}
