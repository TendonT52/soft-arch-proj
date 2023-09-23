"use client";

import { useState } from "react";
import Link from "next/link";
import {
  BriefcaseIcon,
  CheckCircleIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
  KeyIcon,
  MailIcon,
} from "lucide-react";
import { Logo } from "./logo";
import { Button } from "./ui/button";
import { Checkbox } from "./ui/checkbox";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";
import { Separator } from "./ui/separator";
import { Textarea } from "./ui/textarea";

/* DUMMY */
const categories = [
  "E-commerce",
  "Financial Services",
  "Energy and Utilities",
  "Manufacturing and Heavy Industry",
  "Entertainment and Media",
];
/* DUMMY */

const RegisterCompanyForm = () => {
  const [page, setPage] = useState(0);

  return (
    <div className="relative flex flex-1 items-center justify-center">
      <div className="mx-auto flex w-full max-w-lg flex-col items-center justify-center p-8 lg:max-w-sm lg:p-0">
        <div className="mb-6 flex w-full flex-col items-center gap-6">
          <Logo />
          <h1 className="max-w-[85%] text-center text-2xl font-semibold leading-none tracking-tight sm:max-w-none">
            Create a company account
          </h1>
          <p className="max-w-[85%] text-center text-muted-foreground sm:max-w-none">
            {page === 0
              ? "Enter your company information to continue"
              : "Enter your credentials to create your account"}
          </p>
        </div>
        {page === 0 ? (
          <form
            key="page-0"
            className="flex w-full flex-col items-center gap-4"
          >
            <div className="flex w-full gap-4">
              <Input className="flex-1" placeholder="Company name" />
              <Select>
                <SelectTrigger className="flex-1">
                  <SelectValue placeholder="Category" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    {categories.map((category) => (
                      <SelectItem key={category} value={category}>
                        {category}
                      </SelectItem>
                    ))}
                  </SelectGroup>
                </SelectContent>
              </Select>
            </div>
            <div className="flex w-full gap-4">
              <Input className="flex-[1.5]" placeholder="Location" />
              <Input
                className="flex-1 [appearance:textfield] [&::-webkit-inner-spin-button]:appearance-none [&::-webkit-outer-spin-button]:appearance-none"
                placeholder="Phone"
                type="number"
              />
            </div>
            <div className="mb-12 flex w-full">
              <Textarea className="resize-none" placeholder="Description" />
            </div>
            <div className="flex flex-col items-center gap-4">
              <Separator className="w-full" />
              <p className="text-center text-sm text-muted-foreground">
                Already have an account?&nbsp;
                <Link
                  className="font-medium text-primary underline-offset-4 hover:underline"
                  href="/login"
                >
                  Log in
                </Link>
              </p>
            </div>
          </form>
        ) : (
          <form
            key="page-1"
            className="flex w-full flex-col items-center pb-[2.3125rem]"
          >
            <div className="relative mb-4 flex w-full">
              <Label
                className="absolute flex h-full w-10 items-center justify-center"
                htmlFor="email"
              >
                <MailIcon className="h-4 w-4 opacity-50" />
              </Label>
              <Input className="flex-1 pl-10" placeholder="Email" id="email" />
            </div>
            <div className="relative mb-4 flex w-full">
              <Label
                className="absolute flex h-full w-10 items-center justify-center"
                htmlFor="password"
              >
                <KeyIcon className="h-4 w-4 opacity-50" />
              </Label>
              <Input
                className="flex-1 pl-10"
                placeholder="Password"
                type="password"
                id="password"
              />
            </div>
            <div className="relative mb-6 flex w-full">
              <Label
                className="absolute flex h-full w-10 items-center justify-center"
                htmlFor="confirm"
              >
                <CheckCircleIcon className="h-4 w-4 opacity-50" />
              </Label>
              <Input
                className="flex-1 pl-10"
                placeholder="Confirm password"
                type="password"
                id="confirm"
              />
            </div>
            <div className="mb-6 flex items-center gap-2">
              <Checkbox id="terms" />
              <Label htmlFor="terms">Accept terms and conditions</Label>
            </div>
            <div className="flex items-center">
              <Button type="button">
                <BriefcaseIcon className="mr-2 h-4 w-4" />
                Create account
              </Button>
            </div>
          </form>
        )}
      </div>
      {page === 0 ? (
        <Button
          className="absolute bottom-4 right-4 lg:bottom-12 lg:right-12"
          variant="ghost"
          onClick={() => void setPage(1)}
        >
          Next
          <ChevronRightIcon className="ml-2 h-4 w-4 opacity-50" />
        </Button>
      ) : (
        <Button
          className="absolute bottom-4 left-4 lg:bottom-12 lg:left-12"
          variant="ghost"
          onClick={() => void setPage(0)}
        >
          <ChevronLeftIcon className="mr-2 h-4 w-4 opacity-50" />
          Back
        </Button>
      )}
    </div>
  );
};

export { RegisterCompanyForm };
