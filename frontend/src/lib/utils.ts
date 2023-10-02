import { cx } from "cva";
import { type ClassValue } from "cva/types";
import { twMerge } from "tailwind-merge";

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

export function getSearchArray(searchParam: string | string[] | undefined) {
  return searchParam
    ? Array.isArray(searchParam)
      ? searchParam
      : [searchParam]
    : [];
}
