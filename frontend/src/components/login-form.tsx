"use client";

import { KeyIcon, LogIn, MailIcon } from "lucide-react";
import { Logo } from "./logo";
import { SignUpOptionMenu } from "./sign-up-option-menu";
import { Button } from "./ui/button";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Separator } from "./ui/separator";

const LogInForm = () => {
  return (
    <div className="relative flex flex-1 items-center justify-center">
      <div className="mx-auto flex w-full max-w-lg flex-col items-center justify-center p-8 lg:max-w-sm lg:p-0">
        <div className="mb-6 flex w-full flex-col items-center gap-6">
          <Logo />
          <h1 className="max-w-[85%] text-center text-2xl font-semibold leading-none tracking-tight sm:max-w-none">
            <span className="px-2">Sign in to</span>
            <span className="font-bold">InternWise</span>
            <span className="text-primary">Hub</span>
          </h1>
          <p className="max-w-[85%] text-center text-muted-foreground sm:max-w-none">
            Enter your Email and password to sign in
          </p>
        </div>
        <form className="flex w-full flex-col items-center gap-3">
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
          <div className="flex items-center">
            <Button type="button">
              <LogIn className="mr-2 h-4 w-4" />
              Sign In
            </Button>
          </div>
          <div className="flex flex-col items-center gap-4">
            <Separator className="w-full" />
            <p className="text-center text-sm text-muted-foreground">
              Do not have an InternWiseHub ID?&nbsp;
              <SignUpOptionMenu>
                <span className="font-medium text-primary underline-offset-4 hover:underline">
                  Sign Up
                </span>
              </SignUpOptionMenu>
            </p>
          </div>
        </form>
      </div>
    </div>
  );
};

export { LogInForm };
