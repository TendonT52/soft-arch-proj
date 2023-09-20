import Image from "next/image";
import Link from "next/link";
import { Logo } from "./logo";
import { Button } from "./ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";

const Header = () => {
  return (
    <header className="sticky left-0 right-0 top-0 z-50 bg-background/70 backdrop-blur-lg backdrop-saturate-150">
      <div className="container flex h-16 select-none items-center justify-between">
        <Link className="flex font-bold" href="/">
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
          <DropdownMenu modal={false}>
            <DropdownMenuTrigger asChild>
              <Button variant="outline" size="sm">
                Sign up
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end" className="w-56">
              <Link href="/register/student">
                <DropdownMenuItem className="cursor-pointer flex-col items-start gap-2 p-4">
                  <div className="text-base font-medium leading-none tracking-tight">
                    Student
                  </div>
                  <div className="tracking-tight text-muted-foreground">
                    Create a student account
                  </div>
                  <Image
                    className="relative mx-auto h-32 w-32 object-contain"
                    src="/images/sign-up-student.png"
                    alt="Student"
                    width={507}
                    height={515}
                  />
                </DropdownMenuItem>
              </Link>
              <DropdownMenuSeparator />
              <Link href="/register/company">
                <DropdownMenuItem className="cursor-pointer flex-col items-start gap-2 p-4">
                  <div className="text-base font-medium leading-none tracking-tight">
                    Student
                  </div>
                  <div className="tracking-tight text-muted-foreground">
                    Create a student account
                  </div>
                  <Image
                    className="relative mx-auto h-32 w-32 object-contain"
                    src="/images/sign-up-company.png"
                    alt="Company"
                    width={403}
                    height={514}
                  />
                </DropdownMenuItem>
              </Link>
            </DropdownMenuContent>
          </DropdownMenu>
        </div>
      </div>
    </header>
  );
};

export { Header };
