import * as React from "react";
import { CatIcon, SproutIcon } from "lucide-react";
import { cn } from "@/lib/utils";

type LogoProps = React.HTMLAttributes<HTMLDivElement>;

const Logo = React.forwardRef<HTMLDivElement, LogoProps>(
  ({ className, ...props }, ref) => (
    <div
      className={cn("relative inline-flex justify-center", className)}
      ref={ref}
      {...props}
    >
      <CatIcon className="h-6 w-6 fill-background text-foreground" />
      <SproutIcon className="absolute -top-2 h-3.5 w-3.5 text-primary" />
      <div className="absolute inset-0 -z-10 h-6 w-6 rounded-full bg-foreground"></div>
    </div>
  )
);
Logo.displayName = "Logo";

export { Logo };
