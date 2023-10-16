"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { createCompany } from "@/actions/create-company";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Building2Icon,
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
import { Separator } from "./ui/separator";
import { Textarea } from "./ui/textarea";
import { useToast } from "./ui/toaster";

const firstPageSchema = z.object({
  name: z.string().min(1, { message: "Name is required" }),
  category: z.string().min(1, { message: "Category is required" }),
  location: z.string().min(1, { message: "Location is required" }),
  phone: z
    .string()
    .min(1, { message: "Phone number is required" })
    .regex(/^\d+$/, { message: "Phone number must be numerical" }),
  description: z.string().min(1, { message: "Description is required" }),
});

const secondPageSchema = z.object({
  email: z
    .string()
    .min(1, { message: "Email is required" })
    .email({ message: "Invalid email" }),
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

const RegisterCompanyForm = () => {
  const router = useRouter();
  const { toast } = useToast();
  const {
    register,
    formState: { isSubmitting, errors },
    handleSubmit,
    setValue,
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

  const { terms } = useWatch({ control });
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

  const onSubmit = async ({ terms, ...data }: FormData) => {
    const response = await createCompany(data);
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

  useEffect(() => {
    void trigger("terms");
  }, [trigger, terms]);

  return (
    <div className="relative flex flex-1 items-center justify-center">
      <form
        className="mx-auto flex w-full max-w-lg flex-col items-center justify-center p-8 lg:max-w-sm lg:p-0"
        onSubmit={(...a) => void handleSubmit(onSubmit)(...a)}
      >
        <div className="mb-6 flex w-full flex-col items-center gap-2">
          <Logo />
          <h1 className="max-w-[85%] text-center text-2xl font-semibold tracking-tight sm:max-w-none">
            Create a company account
          </h1>
          <p className="max-w-[85%] text-center text-sm text-muted-foreground sm:max-w-none">
            {page === 0
              ? "Enter your company information to continue"
              : "Enter your credentials to create your account"}
          </p>
        </div>
        {page === 0 ? (
          <div key="first" className="flex w-full flex-col items-center gap-4">
            <fieldset className="flex w-full items-center gap-4">
              <Input
                {...register("name")}
                className={cn(
                  "flex-1",
                  errors.name &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
                placeholder="Company name"
              />
              <Input
                {...register("category")}
                className={cn(
                  "flex-1",
                  errors.category &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
                placeholder="Category"
              />
              <FormErrorTooltip
                message={
                  errors.name
                    ? errors.name.message
                    : errors.category
                    ? errors.category.message
                    : undefined
                }
              />
            </fieldset>
            <fieldset className="flex w-full items-center gap-4">
              <Input
                {...register("location")}
                className={cn(
                  "flex-[1.5]",
                  errors.location &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
                placeholder="Location"
              />
              <Input
                {...register("phone")}
                className={cn(
                  "flex-1 [appearance:textfield] [&::-webkit-inner-spin-button]:appearance-none [&::-webkit-outer-spin-button]:appearance-none",
                  errors.phone &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
                placeholder="Phone number"
                type="number"
              />
              <FormErrorTooltip
                message={
                  errors.location
                    ? errors.location.message
                    : errors.phone
                    ? errors.phone.message
                    : undefined
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
            key="second"
            className="flex w-full flex-col items-center pb-[2.3125rem]"
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
                placeholder="Company email"
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
                  <Building2Icon className="mr-2 h-4 w-4" />
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

export { RegisterCompanyForm };
