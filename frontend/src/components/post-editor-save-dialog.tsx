"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { updatePost } from "@/actions/update-post";
import { zodResolver } from "@hookform/resolvers/zod";
import { format } from "date-fns";
import { Loader2Icon } from "lucide-react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { type Post } from "@/types/base/post";
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

/* NAIVE SHIT */
function getUserElements(remove: string[], add: string[]) {
  const userElement = [
    ...remove
      .filter((value) => !add.includes(value))
      .map((value) => ({
        action: "REMOVE" as const,
        value,
      })),
    ...add
      .filter((value) => !remove.includes(value))
      .map((value) => ({
        action: "ADD" as const,
        value,
      })),
  ];
  if (userElement.length === 0) return [{ action: "SAME" as const }];
  return userElement;
}

const formDataSchema = z.object({
  openPositions: z.string().trim(),
  requiredSkills: z.string().trim(),
  benefits: z.string().trim(),
  howTo: z.string().trim(),
});

type FormData = z.infer<typeof formDataSchema>;

type PostEditorSaveDialogProps = {
  postId: string;
  post: Post;
  topic: string;
  description: string;
};

const PostEditorSaveDialog = ({
  postId,
  post,
  topic,
  description,
}: PostEditorSaveDialogProps) => {
  const router = useRouter();
  const { toast } = useToast();
  const {
    register,
    formState: { isSubmitting },
    handleSubmit,
  } = useForm<FormData>({
    mode: "onChange",
    resolver: zodResolver(formDataSchema),
    defaultValues: {
      openPositions: post.openPositions.join(" "),
      requiredSkills: post.requiredSkills.join(" "),
      benefits: post.benefits.join(" "),
      howTo: post.howTo,
    },
  });

  // const {} = useWatch({ control });
  const [period, setPeriod] = useState(post?.period);

  const onSubmit = async (data: FormData) => {
    const response = await updatePost(postId, {
      post: {
        topic,
        description,
        period,
        howTo: data.howTo,
        openPositions: getUserElements(
          post.openPositions,
          data.openPositions.split(/\s+/)
        ),
        requiredSkills: getUserElements(
          post.requiredSkills,
          data.requiredSkills.split(/\s+/)
        ),
        benefits: getUserElements(post.benefits, data.benefits.split(/\s+/)),
      },
    });
    if (response.status === "200") {
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
        <Button className="h-10">Save</Button>
      </DialogTrigger>
      <DialogContent>
        <form
          className="flex w-full flex-col gap-4 p-1"
          onSubmit={(...a) => void handleSubmit(onSubmit)(...a)}
        >
          <DialogHeader className="mb-2">
            <DialogTitle>Additional information</DialogTitle>
            <DialogDescription>
              Help students reach your post easily
            </DialogDescription>
          </DialogHeader>
          <div className="flex flex-col gap-6">
            <div className="flex w-full gap-4">
              <div className="flex flex-1 flex-col gap-2">
                <Label
                  className="w-full text-sm font-medium leading-none"
                  htmlFor="topic"
                >
                  Topic
                </Label>
                <Input id="topic" value={topic} readOnly />
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
                  value={period}
                  onDateChange={(date) =>
                    void setPeriod(
                      date?.from
                        ? date.to
                          ? `${format(date.from, "LLL dd, y")} - ${format(
                              date.to,
                              "LLL dd, y"
                            )}`
                          : format(date.from, "LLL dd, y")
                        : ""
                    )
                  }
                />
              </div>
            </div>
            <div className="flex w-full flex-col gap-2">
              <Label
                className="w-full text-sm font-medium leading-none"
                htmlFor="skills"
              >
                Required skills
                <span className="ml-4 font-normal text-muted-foreground">
                  *Space delimited
                </span>
              </Label>
              <Input
                {...register("requiredSkills")}
                id="requiredSkills"
                placeholder="SQL slamming"
              />
            </div>
            <div className="flex w-full flex-col gap-2">
              <Label
                className="w-full text-sm font-medium leading-none"
                htmlFor="positions"
              >
                Open positions
                <span className="ml-4 font-normal text-muted-foreground">
                  *Space delimited
                </span>
              </Label>
              <Input
                {...register("openPositions")}
                id="openPositions"
                placeholder="Top of the world"
              />
            </div>
            <div className="flex w-full flex-col gap-2">
              <Label
                className="w-full text-sm font-medium leading-none"
                htmlFor="benefits"
              >
                Benefits
                <span className="ml-4 font-normal text-muted-foreground">
                  *Space delimited
                </span>
              </Label>
              <Input
                {...register("benefits")}
                id="benefits"
                placeholder="Coffee"
              />
            </div>
            <div className="flex flex-col gap-2">
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

export { PostEditorSaveDialog };
