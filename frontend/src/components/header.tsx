import Link from "next/link";
import { Logo } from "./logo";
import { SignUpOptionMenu } from "./sign-up-option-menu";
import { Button } from "./ui/button";

const Header = () => {
  return (
    <header className="sticky left-0 right-0 top-0 z-50 bg-background/70 backdrop-blur-lg backdrop-saturate-150">
      <div className="container flex h-16 items-center justify-between">
        <Link className="flex select-none font-bold" href="/">
          <Logo className="mr-2" />
          <span>InternWise</span>
          <span className="text-primary">Hub</span>
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
      </div>
    </header>
  );
};

export { Header };
