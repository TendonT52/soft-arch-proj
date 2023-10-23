"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { createPost } from "@/actions/create-post";
import { zodResolver } from "@hookform/resolvers/zod";
import { format } from "date-fns";
import { Loader2Icon, PlusIcon } from "lucide-react";
import { type DateRange } from "react-day-picker";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { initialEditorState } from "@/lib/lexical";
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
  topic: z.string(),
  openPositions: z.string().trim(),
  requiredSkills: z.string().trim(),
  benefits: z.string().trim(),
  howTo: z.string().trim(),
});

type FormData = z.infer<typeof formDataSchema>;

const NewPostDialog = () => {
  const router = useRouter();
  const { toast } = useToast();
  const {
    register,
    formState: { isSubmitting },
    handleSubmit,
  } = useForm<FormData>({
    mode: "onChange",
    resolver: zodResolver(formDataSchema),
  });

  const [date, setDate] = useState<DateRange>();
  const period = date?.from
    ? date.to
      ? `${format(date.from, "LLL dd, y")} - ${format(date.to, "LLL dd, y")}`
      : format(date.from, "LLL dd, y")
    : "";

  const onSubmit = async (data: FormData) => {
    const regex = /\s+/;
    const response = await createPost({
      post: {
        topic: data.topic,
        description: initialEditorState,
        period,
        howTo: data.howTo,
        openPositions: data.openPositions.split(regex),
        requiredSkills: data.requiredSkills.split(regex),
        benefits: data.benefits.split(regex),
      },
    });
    if (response.status === "201") {
      toast({
        title: "Success",
        description: response.message,
      });
      router.refresh();
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
                  placeholder="Untitled Post"
                  id="topic"
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
                  date={date}
                  onDateChange={setDate}
                />
              </div>
            </fieldset>
            <fieldset className="flex w-full flex-col gap-2">
              <Label
                className="flex w-full justify-between text-sm font-medium leading-none"
                htmlFor="openPositions"
              >
                <div>Open positions</div>
                <span className="ml-4 font-normal text-muted-foreground">
                  Space delimited
                </span>
              </Label>
              <Input
                {...register("openPositions")}
                id="openPositions"
                placeholder="Top of the world"
              />
            </fieldset>
            <fieldset className="flex w-full flex-col gap-2">
              <Label
                className="flex w-full justify-between text-sm font-medium leading-none"
                htmlFor="requiredSkills"
              >
                <div>Required skills</div>
                <span className="ml-4 font-normal text-muted-foreground">
                  Space delimited
                </span>
              </Label>
              <Input
                {...register("requiredSkills")}
                id="requiredSkills"
                placeholder="SQL slamming"
              />
            </fieldset>
            <fieldset className="flex w-full flex-col gap-2">
              <Label
                className="flex w-full justify-between text-sm font-medium leading-none"
                htmlFor="benefits"
              >
                <div>Benefits</div>
                <span className="ml-4 font-normal text-muted-foreground">
                  Space delimited
                </span>
              </Label>
              <Input
                {...register("benefits")}
                id="benefits"
                placeholder="Coffee"
              />
            </fieldset>
            <fieldset className="flex flex-col gap-2">
              <Label
                className="text-sm font-medium leading-none"
                htmlFor="howTo"
              >
                How to apply
              </Label>
              <Textarea
                {...register("howTo")}
                id="howTo"
                className="resize-none"
                placeholder="Run to the office like the flash ⚡"
              />
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

export { NewPostDialog };
