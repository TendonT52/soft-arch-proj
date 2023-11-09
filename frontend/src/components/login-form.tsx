"use client";

import { useRouter } from "next/navigation";
import { zodResolver } from "@hookform/resolvers/zod";
import { KeyIcon, Loader2Icon, LogInIcon, MailIcon } from "lucide-react";
import { signIn } from "next-auth/react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { cn } from "@/lib/utils";
import { FormErrorTooltip } from "./form-error-tooltip";
import { Logo } from "./logo";
import { SignUpOptionMenu } from "./sign-up-option-menu";
import { Button } from "./ui/button";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Separator } from "./ui/separator";
import { useToast } from "./ui/toaster";

const formDataSchema = z.object({
  email: z
    .string()
    .min(1, { message: "Email is required" })
    .email({ message: "Invalid email" }),
  password: z.string().min(1, { message: "Password is required" }),
});

type FormData = z.infer<typeof formDataSchema>;

const LogInForm = () => {
  const router = useRouter();
  const { toast } = useToast();
  const {
    register,
    formState: { isSubmitting, errors },
    handleSubmit,
  } = useForm<FormData>({
    mode: "onChange",
    resolver: zodResolver(formDataSchema),
  });

  const onSubmit = async (data: FormData) => {
    const { email, password } = data;
    const signInResult = await signIn("credentials", {
      email,
      password,
      redirect: false,
    });
    if (signInResult?.error) {
      toast({
        title: "Error",
        description: "Login failed.",
        variant: "destructive",
      });
    } else {
      router.refresh();
      router.push("/dashboard");
    }
  };

  return (
    <div className="relative flex flex-1 items-center justify-center">
      <form
        className="mx-auto flex w-full max-w-lg flex-col items-center justify-center p-8 lg:max-w-sm lg:p-0"
        onSubmit={(...a) => void handleSubmit(onSubmit)(...a)}
      >
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
        <div className="flex w-full flex-col items-center gap-3">
          <fieldset className="relative mb-4 flex w-full items-center gap-4">
            <Label
              className="absolute flex h-full w-10 items-center justify-center"
              htmlFor="email"
            >
              <MailIcon className="h-4 w-4 opacity-50" />
            </Label>
            <Input
              {...register("email")}
              className={cn(
                "flex-1 pl-10",
                errors.email &&
                  "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
              )}
              placeholder="Email"
              id="email"
            />
            <FormErrorTooltip message={errors.email?.message} />
          </fieldset>
          <fieldset className="relative mb-4 flex w-full items-center gap-4">
            <Label
              className="absolute flex h-full w-10 items-center justify-center"
              htmlFor="password"
            >
              <KeyIcon className="h-4 w-4 opacity-50" />
            </Label>
            <Input
              {...register("password")}
              className={cn(
                "flex-1 pl-10",
                errors.password &&
                  "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
              )}
              placeholder="Password"
              type="password"
              id="password"
            />
            <FormErrorTooltip message={errors.password?.message} />
          </fieldset>
          <div className="flex items-center">
            <Button disabled={isSubmitting} type="submit">
              {isSubmitting ? (
                <Loader2Icon className="mr-2 h-4 w-4 animate-spin" />
              ) : (
                <LogInIcon className="mr-2 h-4 w-4" />
              )}
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
        </div>
      </form>
    </div>
  );
};

export { LogInForm };
