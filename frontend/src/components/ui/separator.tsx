"use client";

import * as React from "react";
import { cn } from "@/lib/utils";

export interface SeparatorProps extends React.ComponentPropsWithoutRef<"div"> {
  orientation?: "horizontal" | "vertical";
}

const Separator = React.forwardRef<HTMLDivElement, SeparatorProps>(
  ({ className, orientation = "horizontal", ...props }, ref) => (
    <div
      ref={ref}
      className={cn(
        orientation === "horizontal" ? "h-px border-b" : "w-px border-l",
        className
      )}
      {...props}
    />
  )
);
Separator.displayName = "Separator";

export { Separator };
