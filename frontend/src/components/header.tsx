import * as React from "react";
import Link from "next/link";
import { SearchIcon } from "lucide-react";
import { getServerSession } from "@/lib/auth";
import { Logo } from "./logo";
import { SignUpOptionMenu } from "./sign-up-option-menu";
import { Button } from "./ui/button";
import { UserAccountNav } from "./user-account-nav";

const Header = async () => {
  const session = await getServerSession();
  const user = session?.user;

  return (
    <header className="container sticky left-0 right-0 top-0 z-50 flex h-16 items-center justify-between bg-background/70 backdrop-blur-xl backdrop-saturate-150">
      <Link className="flex font-bold" href="/">
        <Logo className="mr-2" />
        <div>
          InternWise
          <span className="text-primary">Hub</span>
        </div>
      </Link>
      {user ? (
        <div className="flex items-center gap-4 text-sm">
          <Button
            className="text-muted-foreground"
            variant="outline"
            size="sm"
            asChild
          >
            <Link href="/posts">
              <SearchIcon className="mr-2 h-4 w-4 shrink-0" />
              Search posts
            </Link>
          </Button>
          <UserAccountNav user={user} />
        </div>
      ) : (
        <div className="flex items-center gap-4 text-sm font-medium">
          <Link
            className="text-foreground transition-colors hover:text-foreground/90"
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
      )}
    </header>
  );
};

export { Header };
