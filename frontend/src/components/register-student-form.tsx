"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { createStudent } from "@/actions/create-student";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  BriefcaseIcon,
  CheckCircleIcon,
  ChevronLeftIcon,
  ChevronRightIcon,
  KeyIcon,
  Loader2Icon,
  MailIcon,
} from "lucide-react";
import { useForm, useWatch } from "react-hook-form";
import { z } from "zod";
import { cn } from "@/lib/utils";
import { FormErrorTooltip } from "./form-error-tooltip";
import { Logo } from "./logo";
import { Button } from "./ui/button";
import { Checkbox } from "./ui/checkbox";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger } from "./ui/select";
import { Separator } from "./ui/separator";
import { Textarea } from "./ui/textarea";
import { useToast } from "./ui/toaster";

const years = [1, 2, 3, 4];

const firstPageSchema = z.object({
  firstName: z
    .string()
    .min(1, { message: "First name is required" })
    .regex(/^[a-zA-Z]+$/, { message: "First name must be alphabetical" }),
  lastName: z
    .string()
    .min(1, { message: "Last name is required" })
    .regex(/^[a-zA-Z]+$/, { message: "Last name must be alphabetical" }),
  faculty: z
    .string()
    .min(1, { message: "Faculty is required" })
    .regex(/^[a-zA-Z]+$/, { message: "Faculty must be alphabetical" }),
  major: z
    .string()
    .min(1, { message: "Major is required" })
    .regex(/^[a-zA-Z]+$/, { message: "Major must be alphabetical" }),
  year: z.number({ required_error: "Year is required" }),
  description: z.string().min(1, { message: "Description is required" }),
});

const secondPageSchema = z.object({
  email: z
    .string()
    .min(1, { message: "Email is required" })
    .regex(/^\d{10}@student\.chula\.ac\.th$/, {
      message: "Email must be studentID with @student.chula.ac.th",
    }),
  password: z
    .string()
    .min(1, { message: "Password is required" })
    .min(6, { message: "Password length must be greater than 6" }),
  passwordConfirm: z
    .string()
    .min(1, { message: "Please confirm your password" }),
  terms: z.boolean(),
});

const formDataSchema = firstPageSchema
  .merge(secondPageSchema)
  .refine((data) => data.password === data.passwordConfirm, {
    message: "Passwords do not match",
    path: ["passwordConfirm"],
  })
  .refine((data) => data.terms, {
    message: "ACCEPT IT NOW!",
    path: ["terms"],
  });

type FormData = z.infer<typeof formDataSchema>;

const RegisterStudentForm = () => {
  const router = useRouter();
  const { toast } = useToast();
  const {
    register,
    formState: { isSubmitting, errors },
    handleSubmit,
    setValue,
    clearErrors,
    trigger,
    control,
  } = useForm<FormData>({
    mode: "onChange",
    resolver: zodResolver(formDataSchema),
    defaultValues: {
      passwordConfirm: "",
      terms: false,
    },
  });

  const { year, terms } = useWatch({ control });
  const [page, setPage] = useState(0);

  const goPreviousPage = () => void setPage(0);

  const goNextPage = async () => {
    const validationResult = await trigger(firstPageSchema.keyof().options, {
      shouldFocus: true,
    });
    if (validationResult) {
      setPage(1);
    }
  };

  const onSubmit = async ({
    terms,
    firstName,
    lastName,
    ...data
  }: FormData) => {
    const response = await createStudent({
      ...data,
      name: `${firstName} ${lastName}`,
    });
    if (response.status === "201") {
      toast({
        title: "Success",
        description: response.message,
      });
      router.push("/");
    } else {
      toast({
        title: "Error",
        description: response.message,
        variant: "destructive",
      });
    }
  };

  useEffect(() => void trigger("terms"), [trigger, terms]);

  return (
    <div className="relative flex flex-1 items-center justify-center">
      <form
        className="mx-auto flex w-full max-w-lg flex-col items-center justify-center p-8 lg:max-w-sm lg:p-0"
        onSubmit={(...a) => void handleSubmit(onSubmit)(...a)}
      >
        <div className="mb-6 flex w-full flex-col items-center gap-2">
          <Logo />
          <h1 className="max-w-[85%] text-center text-2xl font-semibold tracking-tight sm:max-w-none">
            Create a student account
          </h1>
          <p className="max-w-[85%] text-center text-sm text-muted-foreground sm:max-w-none">
            {page === 0
              ? "Enter your personal information to continue"
              : "Enter your credentials to create your account"}
          </p>
        </div>
        {page === 0 ? (
          <div key="first" className="flex w-full flex-col items-center gap-4">
            <fieldset className="flex w-full items-center gap-4">
              <Input
                {...register("firstName")}
                className={cn(
                  "flex-1",
                  errors.firstName &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
                placeholder="First name"
              />
              <Input
                {...register("lastName")}
                className={cn(
                  "flex-1",
                  errors.lastName &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
                placeholder="Last name"
              />
              <FormErrorTooltip
                message={errors.firstName?.message ?? errors.lastName?.message}
              />
            </fieldset>
            <fieldset className="relative flex w-full items-center gap-4">
              <Input
                {...register("faculty")}
                className={cn(
                  "flex-1",
                  errors.faculty &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
                placeholder="Faculty"
              />
              <Input
                {...register("major")}
                className={cn(
                  "flex-1",
                  errors.major &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
                placeholder="Major"
              />
              <Select
                value={year?.toString()}
                onValueChange={(year) => {
                  setValue("year", parseInt(year));
                  clearErrors("year");
                }}
              >
                <SelectTrigger
                  className={cn(
                    "flex-[0.5]",
                    year === undefined && "text-muted-foreground",
                    errors.year &&
                      "ring-2 !ring-destructive ring-offset-2 focus-visible:ring-destructive"
                  )}
                >
                  {year ?? "Year"}
                </SelectTrigger>
                <SelectContent>
                  {years.map((year) => (
                    <SelectItem
                      className="pr-0"
                      key={year}
                      value={year.toString()}
                    >
                      {year}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              <FormErrorTooltip
                message={
                  errors.faculty?.message ??
                  errors.major?.message ??
                  errors.year?.message
                }
              />
            </fieldset>
            <fieldset className="mb-12 flex w-full items-center gap-4">
              <Textarea
                {...register("description")}
                className={cn(
                  "flex-1 resize-none",
                  errors.description &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
                placeholder="Description"
              />
              <FormErrorTooltip message={errors.description?.message} />
            </fieldset>
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
        ) : (
          <div
            className="flex w-full flex-col items-center pb-[2.3125rem]"
            key="second"
          >
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
                id="email"
                placeholder="Chula email"
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
                id="password"
                placeholder="Password"
                type="password"
              />
              <FormErrorTooltip message={errors.password?.message} />
            </fieldset>
            <fieldset className="relative mb-6 flex w-full items-center gap-4">
              <Label
                className="absolute flex h-full w-10 items-center justify-center"
                htmlFor="passwordConfirm"
              >
                <CheckCircleIcon className="h-4 w-4 opacity-50" />
              </Label>
              <Input
                {...register("passwordConfirm")}
                className={cn(
                  "flex-1 pl-10",
                  errors.passwordConfirm &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
                id="passwordConfirm"
                placeholder="Confirm password"
                type="password"
              />
              <FormErrorTooltip message={errors.passwordConfirm?.message} />
            </fieldset>
            <fieldset className="mb-5 flex h-5 items-center">
              <Checkbox
                className="mr-2"
                id="terms"
                checked={terms}
                onCheckedChange={(checked: boolean) => {
                  setValue("terms", checked);
                }}
              />
              <Label className="mr-4 flex h-5 items-center" htmlFor="terms">
                Accept the terms and conditions
              </Label>
              <FormErrorTooltip message={errors.terms?.message} />
            </fieldset>
            <div className="flex items-center">
              <Button disabled={isSubmitting} type="submit">
                {isSubmitting ? (
                  <Loader2Icon className="mr-2 h-4 w-4 animate-spin" />
                ) : (
                  <BriefcaseIcon className="mr-2 h-4 w-4" />
                )}
                Create account
              </Button>
            </div>
          </div>
        )}
      </form>
      {page === 0 ? (
        <Button
          className="absolute bottom-4 right-4 lg:bottom-12 lg:right-12"
          variant="ghost"
          type="button"
          onClick={() => void goNextPage()}
        >
          Next
          <ChevronRightIcon className="ml-2 h-4 w-4 opacity-50" />
        </Button>
      ) : (
        <Button
          className="absolute bottom-4 left-4 lg:bottom-12 lg:left-12"
          variant="ghost"
          type="button"
          onClick={goPreviousPage}
        >
          <ChevronLeftIcon className="mr-2 h-4 w-4 opacity-50" />
          Back
        </Button>
      )}
    </div>
  );
};

export { RegisterStudentForm };
