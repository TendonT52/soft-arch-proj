import * as React from "react";
import { CatIcon, SproutIcon } from "lucide-react";
import { cn } from "@/lib/utils";

interface LogoProps extends React.HTMLAttributes<HTMLDivElement> {}

const Logo = React.forwardRef<HTMLDivElement, LogoProps>(
  ({ className, ...props }, ref) => (
    <div
      className={cn("relative flex justify-center", className)}
      ref={ref}
      {...props}
    >
      <CatIcon className="h-9 w-9 fill-background text-foreground" />
      <SproutIcon className="absolute -top-2.5 h-5 w-5 text-primary" />
      <div className="absolute inset-0 -z-10 h-9 w-9 rounded-full bg-foreground"></div>
    </div>
  )
);
Logo.displayName = "Logo";

export { Logo };
