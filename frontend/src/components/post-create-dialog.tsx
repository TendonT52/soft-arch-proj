"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { createPost } from "@/actions/create-post";
import { zodResolver } from "@hookform/resolvers/zod";
import { Loader2Icon, PlusIcon } from "lucide-react";
import { type DateRange } from "react-day-picker";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { initialEditorState } from "@/lib/lexical";
import { cn, formatPeriod } from "@/lib/utils";
import { DatePickerWithRange } from "./date-range-picker";
import { Button } from "./ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "./ui/dialog";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Textarea } from "./ui/textarea";
import { useToast } from "./ui/toaster";

const formDataSchema = z.object({
  topic: z.string().min(1, { message: "Topic is required" }),
  openPositions: z
    .string()
    .trim()
    .min(1, { message: "At least 1 open position is required" }),
  requiredSkills: z
    .string()
    .trim()
    .min(1, { message: "At least 1 required skill is required" }),
  benefits: z
    .string()
    .trim()
    .min(1, { message: "At least 1 benefit is required" }),
  howTo: z.string().trim().min(1, { message: "How to is required" }),
});

type FormData = z.infer<typeof formDataSchema>;

const PostCreateDialog = () => {
  const router = useRouter();
  const { toast } = useToast();

  const {
    register,
    formState: { isSubmitting, errors },
    handleSubmit,
  } = useForm<FormData>({
    mode: "onChange",
    shouldUseNativeValidation: true,
    resolver: zodResolver(formDataSchema),
  });

  const [date, setDate] = useState<DateRange>();

  const onSubmit = async (data: FormData) => {
    const regex = /\s+/;
    const response = await createPost({
      post: {
        topic: data.topic,
        description: initialEditorState,
        period: formatPeriod(date),
        howTo: data.howTo,
        openPositions: [...new Set(data.openPositions.split(regex))],
        requiredSkills: [...new Set(data.requiredSkills.split(regex))],
        benefits: [...new Set(data.benefits.split(regex))],
      },
    });
    if (response.status === "201") {
      toast({
        title: "Success",
        description: response.message,
      });
      router.refresh();
      router.push(`/editor/${response.id}`);
    } else {
      toast({
        title: "Error",
        description: response.message,
        variant: "destructive",
      });
    }
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className="h-10">
          <PlusIcon className="mr-2 h-4 w-4" />
          New post
        </Button>
      </DialogTrigger>
      <DialogContent>
        <form
          className="flex w-full flex-col gap-4 p-1"
          onSubmit={(...a) => void handleSubmit(onSubmit)(...a)}
        >
          <DialogHeader className="mb-2">
            <DialogTitle>Create new post</DialogTitle>
            <DialogDescription>
              Enter information to create your post
            </DialogDescription>
          </DialogHeader>
          <div className="flex flex-col gap-6">
            <fieldset className="flex w-full gap-4">
              <div className="flex flex-1 flex-col gap-2">
                <Label
                  className="w-full text-sm font-medium leading-none"
                  htmlFor="topic"
                >
                  Topic
                </Label>
                <Input
                  {...register("topic")}
                  id="topic"
                  className={cn(
                    errors.topic &&
                      "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                  )}
                  placeholder="Untitled Post"
                />
              </div>
              <div className="flex flex-1 flex-col gap-2">
                <Label
                  htmlFor="period"
                  className="w-full text-sm font-medium leading-none"
                >
                  Period
                </Label>
                <DatePickerWithRange
                  id="period"
                  className={cn(
                    !date &&
                      "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                  )}
                  date={date}
                  onDateChange={setDate}
                />
              </div>
            </fieldset>
            <fieldset className="flex w-full gap-4">
              <div className="flex w-full flex-col gap-2">
                <Label
                  className="flex w-full justify-between text-sm font-medium leading-none"
                  htmlFor="openPositions"
                >
                  Open positions
                  <span className="ml-4 font-normal text-muted-foreground">
                    Space delimited
                  </span>
                </Label>
                <Input
                  {...register("openPositions")}
                  id="openPositions"
                  className={cn(
                    errors.openPositions &&
                      "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                  )}
                  placeholder="Top of the world"
                />
              </div>
            </fieldset>
            <fieldset className="flex w-full gap-4">
              <div className="flex w-full flex-col gap-2">
                <Label
                  className="flex w-full justify-between text-sm font-medium leading-none"
                  htmlFor="requiredSkills"
                >
                  Required skills
                  <span className="ml-4 font-normal text-muted-foreground">
                    Space delimited
                  </span>
                </Label>
                <Input
                  {...register("requiredSkills")}
                  id="requiredSkills"
                  className={cn(
                    errors.requiredSkills &&
                      "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                  )}
                  placeholder="SQL slamming"
                />
              </div>
            </fieldset>
            <fieldset className="flex w-full gap-4">
              <div className="flex w-full flex-col gap-2">
                <Label
                  className="flex w-full justify-between text-sm font-medium leading-none"
                  htmlFor="benefits"
                >
                  Benefits
                  <span className="ml-4 font-normal text-muted-foreground">
                    Space delimited
                  </span>
                </Label>
                <Input
                  {...register("benefits")}
                  id="benefits"
                  className={cn(
                    errors.benefits &&
                      "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                  )}
                  placeholder="Coffee"
                />
              </div>
            </fieldset>
            <fieldset className="flex w-full gap-4">
              <div className="flex w-full flex-col gap-2">
                <Label
                  className="text-sm font-medium leading-none"
                  htmlFor="howTo"
                >
                  How to apply
                </Label>
                <Textarea
                  {...register("howTo")}
                  id="howTo"
                  className={cn(
                    "resize-none",
                    errors.howTo &&
                      "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                  )}
                  placeholder="Run to the office like the flash ⚡"
                />
              </div>
            </fieldset>
          </div>
          <DialogFooter className="mt-2 flex sm:justify-center">
            <Button size="sm" disabled={isSubmitting} type="submit">
              {isSubmitting && (
                <Loader2Icon className="mr-2 h-4 w-4 animate-spin" />
              )}
              Confirm
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
};

export { PostCreateDialog };
