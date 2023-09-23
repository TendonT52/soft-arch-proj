import * as React from "react";
import Link from "next/link";
import { cn } from "@/lib/utils";
import { Logo } from "./logo";
import { SignUpOptionMenu } from "./sign-up-option-menu";
import { Button } from "./ui/button";

interface HeaderProps extends React.HTMLAttributes<HTMLDivElement> {}

const Header = React.forwardRef<HTMLDivElement, HeaderProps>(
  ({ className, ...props }, ref) => {
    return (
      <header
        className={cn(
          "container sticky left-0 right-0 top-0 z-50 flex h-16 items-center justify-between bg-background/70 backdrop-blur-lg backdrop-saturate-150",
          className
        )}
        ref={ref}
        {...props}
      >
        <Link className="flex font-bold" href="/">
          <Logo className="mr-2" />
          <div>
            InternWise
            <span className="text-primary">Hub</span>
          </div>
        </Link>
        <div className="flex items-center gap-4 text-sm font-medium">
          <Link
            className="text-foreground transition-colors hover:text-primary"
            href="/login"
          >
            Login
          </Link>
          <SignUpOptionMenu align="end">
            <Button variant="outline" size="sm">
              Sign up
            </Button>
          </SignUpOptionMenu>
        </div>
      </header>
    );
  }
);
Header.displayName = "Header";

export { Header };
