import Link from "next/link";
import { ChevronLeftIcon } from "lucide-react";
import { Button } from "@/components/ui/button";
import { ModeToggle } from "@/components/mode-toggle";

type LayoutProps = {
  children: React.ReactNode;
};

export default function Layout({ children }: LayoutProps) {
  return (
    <div className="relative flex min-h-screen items-start md:container">
      <div className="sticky top-0 hidden h-screen flex-1 flex-col items-start py-6 lg:flex">
        <Button variant="ghost" asChild>
          <Link href="/dashboard/posts">
            <ChevronLeftIcon className="mr-2 h-4 w-4" />
            Back
          </Link>
        </Button>
      </div>
      {children}
      <div className="sticky top-0 hidden h-screen flex-1 flex-col items-end justify-end py-6 lg:flex">
        <ModeToggle />
      </div>
    </div>
  );
}
