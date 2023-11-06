"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { createPost } from "@/actions/create-post";
import { zodResolver } from "@hookform/resolvers/zod";
import { Loader2Icon, PlusIcon } from "lucide-react";
import { type DateRange } from "react-day-picker";
import { useForm, useWatch } from "react-hook-form";
import { z } from "zod";
import { PostField } from "@/types/base/post";
import { initialEditorState } from "@/lib/lexical";
import { cn, formatPeriod } from "@/lib/utils";
import { DatePickerWithRange } from "./date-range-picker";
import { FormErrorTooltip } from "./form-error-tooltip";
import { PostFieldInput } from "./post-field-input";
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
    .array(z.string())
    .min(1, { message: "At least 1 open position is required" }),
  requiredSkills: z
    .array(z.string())
    .min(1, { message: "At least 1 required skill is required" }),
  benefits: z
    .array(z.string())
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
    trigger,
    setValue,
    control,
  } = useForm<FormData>({
    mode: "onChange",
    resolver: zodResolver(formDataSchema),
  });

  const { openPositions, requiredSkills, benefits } = useWatch({ control });
  const [date, setDate] = useState<DateRange>();

  const onSubmit = async (data: FormData) => {
    const response = await createPost({
      post: {
        topic: data.topic,
        description: initialEditorState,
        period: formatPeriod(date),
        howTo: data.howTo,
        openPositions: data.openPositions,
        requiredSkills: data.requiredSkills,
        benefits: data.benefits,
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

  useEffect(() => {
    console.log(openPositions);
  }, [openPositions]);

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
          className="flex w-full flex-col gap-4 overflow-auto p-1"
          onSubmit={(...a) => void handleSubmit(onSubmit)(...a)}
        >
          <DialogHeader className="mb-2">
            <DialogTitle>New post</DialogTitle>
            <DialogDescription>
              Enter information to create your post
            </DialogDescription>
          </DialogHeader>
          <div className="flex flex-col items-start gap-6">
            <div className="flex w-full items-center gap-4">
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
              <div className="flex flex-1 flex-col gap-2 overflow-auto">
                <Label
                  htmlFor="period"
                  className="text-sm font-medium leading-none"
                >
                  Period
                </Label>
                <DatePickerWithRange
                  id="period"
                  date={date}
                  onDateChange={setDate}
                />
              </div>
              <FormErrorTooltip
                className="relative top-3"
                message={errors.topic?.message}
              />
            </div>
            <div className="flex w-full gap-4">
              <div className="flex w-full flex-col gap-2">
                <Label
                  className="flex w-full justify-between text-sm font-medium leading-none"
                  htmlFor="openPositions"
                >
                  Open positions
                </Label>
                <div className="flex w-full items-center gap-4">
                  <PostFieldInput
                    id="openPositions"
                    className={cn(
                      "flex-1",
                      errors.openPositions &&
                        "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                    )}
                    field={PostField.openPositions}
                    tags={openPositions}
                    onTagsChange={(tags) => {
                      setValue("openPositions", tags);
                      void trigger("openPositions");
                    }}
                  />
                  <FormErrorTooltip message={errors.openPositions?.message} />
                </div>
              </div>
            </div>
            <div className="flex w-full gap-4">
              <div className="flex w-full flex-col gap-2">
                <Label
                  className="flex w-full justify-between text-sm font-medium leading-none"
                  htmlFor="requiredSkills"
                >
                  Required skills
                </Label>

                <div className="flex w-full items-center gap-4">
                  <PostFieldInput
                    id="requiredSkills"
                    className={cn(
                      "flex-1",
                      errors.requiredSkills &&
                        "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                    )}
                    field={PostField.requiredSkills}
                    tags={requiredSkills}
                    onTagsChange={(tags) => {
                      setValue("requiredSkills", tags);
                      void trigger("requiredSkills");
                    }}
                  />
                  <FormErrorTooltip message={errors.requiredSkills?.message} />
                </div>
              </div>
            </div>
            <div className="flex w-full gap-4">
              <div className="flex w-full flex-col gap-2">
                <Label
                  className="flex w-full justify-between text-sm font-medium leading-none"
                  htmlFor="benefits"
                >
                  Benefits
                </Label>

                <div className="flex w-full items-center gap-4">
                  <PostFieldInput
                    id="benefits"
                    className={cn(
                      "flex-1",
                      errors.benefits &&
                        "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                    )}
                    field={PostField.benefits}
                    tags={benefits}
                    onTagsChange={(tags) => {
                      setValue("benefits", tags);
                      void trigger("benefits");
                    }}
                  />
                  <FormErrorTooltip message={errors.benefits?.message} />
                </div>
              </div>
            </div>
            <div className="flex w-full gap-4">
              <div className="flex w-full flex-col gap-2">
                <Label
                  className="text-sm font-medium leading-none"
                  htmlFor="howTo"
                >
                  How to apply
                </Label>
                <div className="flex items-center gap-4">
                  <Textarea
                    {...register("howTo")}
                    id="howTo"
                    className={cn(
                      "flex-1 resize-none",
                      errors.howTo &&
                        "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                    )}
                    placeholder="Run to the office like the flash ⚡"
                  />
                  <FormErrorTooltip message={errors.howTo?.message} />
                </div>
              </div>
            </div>
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
