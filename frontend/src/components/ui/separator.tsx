"use client";

import { cn } from "@/lib/utils";
import * as React from "react";

export interface SeparatorProps extends React.ComponentPropsWithoutRef<"div"> {
  orientation?: "horizontal" | "vertical";
}

const Separator = React.forwardRef<HTMLDivElement, SeparatorProps>(
  ({ className, orientation = "horizontal", ...props }, ref) => (
    <div
      ref={ref}
      className={cn(
        "bg-border",
        orientation === "horizontal"
          ? "h-0 w-auto border-b"
          : "h-auto w-0 border-l",
        className
      )}
      {...props}
    />
  )
);
Separator.displayName = "Separator";

export { Separator };
