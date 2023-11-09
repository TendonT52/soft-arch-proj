import Image from "next/image";
import Link from "next/link";
import { cn } from "@/lib/utils";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import { Separator } from "./ui/separator";

type SignUpOptionMenuProps = {
  align?: "start" | "center" | "end";
  side?: "left" | "right";
  direction?: "row" | "column";
  children?: JSX.Element;
};

const SignUpOptionMenu = ({
  align,
  side,
  direction = "column",
  children,
}: SignUpOptionMenuProps) => {
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>{children}</DropdownMenuTrigger>
      <DropdownMenuContent
        align={align}
        side={side}
        className={cn(
          "flex",
          direction === "column" ? "w-56 flex-col" : "h-56 flex-row"
        )}
      >
        <Link href="/register/student">
          <DropdownMenuItem className="cursor-pointer flex-col items-start gap-2 p-4">
            <div className="text-base font-medium leading-none tracking-tight">
              Student
            </div>
            <div className="text-muted-foreground">
              Create a student account
            </div>
            <Image
              className="relative mx-auto h-32 w-32 object-contain"
              src="/images/sign-up-student.png"
              alt="Student"
              height={515}
              width={507}
            />
          </DropdownMenuItem>
        </Link>
        <Separator
          className={direction === "column" ? "my-1" : "mx-1"}
          orientation={direction === "column" ? "horizontal" : "vertical"}
        />
        <Link href="/register/company">
          <DropdownMenuItem className="cursor-pointer flex-col items-start gap-2 p-4">
            <div className="text-base font-medium leading-none tracking-tight">
              Company
            </div>
            <div className="text-muted-foreground">
              Create a company account
            </div>
            <Image
              className="relative mx-auto h-32 w-32 object-contain"
              src="/images/sign-up-company.png"
              alt="Company"
              height={514}
              width={403}
            />
          </DropdownMenuItem>
        </Link>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export { SignUpOptionMenu };
