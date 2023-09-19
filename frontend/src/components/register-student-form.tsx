"use client";

import { useState } from "react";
import Link from "next/link";
import {
  BriefcaseIcon,
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

const faculties = ["Engineering", "Science", "Arts", "Business"];
const majors = ["Computer", "Mathematics", "Physics", "Chemistry"];
const years = ["1", "2", "3", "4"];

const RegisterStudentForm = () => {
  const [page, setPage] = useState(0);

  return (
    <div className="relative flex flex-1 items-center justify-center">
      <div className="mx-auto flex w-full max-w-sm flex-col items-center justify-center p-8 lg:p-0">
        <div className="mb-6 flex flex-col items-center gap-6">
          <Logo />
          <h1 className="text-center text-2xl font-semibold tracking-tight">
            Create a student account
          </h1>
          <p className="text-center tracking-tight text-muted-foreground">
            Enter your personal information to continue
          </p>
        </div>
        {page === 0 ? (
          <form key="page-0" className="mb-12 flex flex-col items-center gap-4">
            <div className="flex w-full gap-4">
              <Input className="flex-1" placeholder="First name" />
              <Input className="flex-1" placeholder="Last name" />
            </div>
            <div className="flex w-full gap-4">
              <Select>
                <SelectTrigger className="flex-1">
                  <SelectValue placeholder="Faculty" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    {faculties.map((faculty) => (
                      <SelectItem key={faculty} value={faculty}>
                        {faculty}
                      </SelectItem>
                    ))}
                  </SelectGroup>
                </SelectContent>
              </Select>
              <Select>
                <SelectTrigger className="flex-1">
                  <SelectValue placeholder="Major" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    {majors.map((faculty) => (
                      <SelectItem key={faculty} value={faculty}>
                        {faculty}
                      </SelectItem>
                    ))}
                  </SelectGroup>
                </SelectContent>
              </Select>
              <Select>
                <SelectTrigger className="flex-[0.5]">
                  <SelectValue placeholder="Year" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    {years.map((years) => (
                      <SelectItem key={years} value={years.toString()}>
                        {years}
                      </SelectItem>
                    ))}
                  </SelectGroup>
                </SelectContent>
              </Select>
            </div>
            <div className="flex w-full">
              <Textarea className="resize-none" placeholder="Description" />
            </div>
          </form>
        ) : (
          <form
            key="page-1"
            className="mb-10 flex w-full flex-col items-center"
          >
            <div className="relative mb-4 flex w-full">
              <Label
                className="absolute flex h-full w-10 items-center justify-center"
                htmlFor="email"
              >
                <MailIcon className="h-4 w-4 opacity-50" />
              </Label>
              <Input
                className="flex-1 pl-10"
                placeholder="Chula email"
                id="email"
              />
            </div>
            <div className="relative mb-6 flex w-full">
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

export { RegisterStudentForm };
