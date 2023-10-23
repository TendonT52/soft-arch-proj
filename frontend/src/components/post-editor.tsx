"use client";

import * as React from "react";
import { useEffect, useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { updatePost } from "@/actions/update-post";
import { zodResolver } from "@hookform/resolvers/zod";
import { TRANSFORMERS } from "@lexical/markdown";
import { LexicalComposer } from "@lexical/react/LexicalComposer";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import LexicalErrorBoundary from "@lexical/react/LexicalErrorBoundary";
import { HistoryPlugin } from "@lexical/react/LexicalHistoryPlugin";
import { LinkPlugin } from "@lexical/react/LexicalLinkPlugin";
import { ListPlugin } from "@lexical/react/LexicalListPlugin";
import { MarkdownShortcutPlugin } from "@lexical/react/LexicalMarkdownShortcutPlugin";
import { OnChangePlugin } from "@lexical/react/LexicalOnChangePlugin";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import { ChevronLeftIcon, Loader2Icon } from "lucide-react";
import { type DateRange } from "react-day-picker";
import { useForm } from "react-hook-form";
import TextareaAutosize from "react-textarea-autosize";
import { z } from "zod";
import { type Post } from "@/types/base/post";
import { editorConfig } from "@/lib/lexical";
import { cn, formatPeriod, parsePeriod } from "@/lib/utils";
import { DatePickerWithRange } from "./date-range-picker";
import { CodeHighlightPlugin } from "./lexical/code-highlight-plugin";
import { ListMaxIndentLevelPlugin } from "./lexical/list-max-index-level-plugin";
import { ToolbarPlugin } from "./lexical/toolbar-plugin";
import { Loading } from "./loading";
import { ModeToggle } from "./mode-toggle";
import { Button } from "./ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "./ui/card";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Textarea } from "./ui/textarea";
import { useToast } from "./ui/toaster";

const formDataSchema = z.object({
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

type PostEditorProps = {
  post: Post & {
    postId: string;
  };
};

type FormData = z.infer<typeof formDataSchema>;

/* NAIVE SHIT */
const getUserElements = (remove: string[], add: string[]) => {
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
};

const PostEditor = ({ post }: PostEditorProps) => {
  const editorState = post.description;
  const [description, setDescription] = useState<string>(post.description);
  const [topic, setTopic] = useState(post.topic);
  const [loading, setLoading] = useState(true);

  useEffect(() => void setLoading(false), []);

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
    defaultValues: {
      openPositions: post.openPositions.join(" "),
      requiredSkills: post.requiredSkills.join(" "),
      benefits: post.benefits.join(" "),
      howTo: post.howTo,
    },
  });

  const [date, setDate] = useState<DateRange | undefined>(
    parsePeriod(post.period)
  );
  const period = formatPeriod(date);

  const onSubmit = async (data: FormData) => {
    const regex = /\s+/;
    const response = await updatePost(post.postId, {
      post: {
        topic: topic || "Untitled Post",
        description,
        period,
        howTo: data.howTo,
        openPositions: getUserElements(
          post.openPositions,
          data.openPositions.split(regex)
        ),
        requiredSkills: getUserElements(
          post.requiredSkills,
          data.requiredSkills.split(regex)
        ),
        benefits: getUserElements(post.benefits, data.benefits.split(regex)),
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
    <div className="container relative flex min-h-screen justify-center">
      <div className="sticky top-0 hidden h-screen flex-1 flex-col items-start py-6 lg:flex">
        <Button variant="ghost" asChild>
          <Link href="/dashboard/posts">
            <ChevronLeftIcon className="mr-2 h-4 w-4" />
            Back
          </Link>
        </Button>
      </div>
      <div className="flex w-full max-w-3xl flex-col">
        <LexicalComposer initialConfig={{ ...editorConfig, editorState }}>
          <div className="relative flex flex-col items-start">
            <div className="sticky top-0 z-10 flex w-full justify-center bg-background pt-4">
              <ToolbarPlugin className="rounded-full border bg-background px-4 shadow-sm" />
            </div>
            <TextareaAutosize
              className="flex h-[5.625rem] w-full max-w-none resize-none appearance-none items-end bg-transparent px-8 pb-0.5 pt-12 text-4xl font-extrabold text-[#111827] scrollbar-hide focus:outline-none focus-visible:outline-none dark:text-[#ffffff]"
              value={topic}
              placeholder="Untitled Post"
              spellCheck={false}
              onChange={(e) => void setTopic(e.target.value)}
              aria-label="title"
            />
            {loading ? (
              <div className="relative flex min-h-[32rem] w-full items-center justify-center">
                <Loading />
              </div>
            ) : (
              <div className="relative w-full">
                <RichTextPlugin
                  contentEditable={
                    <ContentEditable
                      className="prose prose-green relative min-h-[32rem] max-w-none flex-1 bg-background px-8 pb-20 pt-8 dark:prose-invert focus:outline-none"
                      spellCheck={false}
                    />
                  }
                  placeholder={
                    <div className="prose prose-green pointer-events-none absolute left-0 right-0 top-0 max-w-none p-8">
                      <p className="text-muted-foreground">
                        Enter post description...
                      </p>
                    </div>
                  }
                  ErrorBoundary={LexicalErrorBoundary}
                />
              </div>
            )}
          </div>
          <HistoryPlugin />
          <CodeHighlightPlugin />
          <ListPlugin />
          <LinkPlugin />
          <ListMaxIndentLevelPlugin maxDepth={1} />
          <MarkdownShortcutPlugin transformers={TRANSFORMERS} />
          <OnChangePlugin
            onChange={(editorState) => {
              setDescription(JSON.stringify(editorState.toJSON()));
              console.log(JSON.stringify(editorState.toJSON()));
            }}
          />
        </LexicalComposer>
        <Card className="mx-auto w-full max-w-3xl rounded-none border-transparent border-t-border bg-background">
          <CardHeader className="p-8">
            <CardTitle>Additional information</CardTitle>
            <CardDescription>
              Help students reach your post easily
            </CardDescription>
          </CardHeader>
          <CardContent className="flex flex-col gap-6 p-8 pt-0">
            <fieldset className="flex w-full gap-4">
              <div className="flex flex-1 flex-col gap-2">
                <Label
                  className="w-full text-sm font-medium leading-none"
                  htmlFor="topic"
                >
                  Topic
                </Label>
                <Input id="topic" value={topic || "Untitled Post"} readOnly />
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
                className="w-full text-sm font-medium leading-none"
                htmlFor="openPositions"
              >
                Open positions
                <span className="ml-4 font-normal text-muted-foreground">
                  *Space delimited
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
            </fieldset>
            <fieldset className="flex w-full flex-col gap-2">
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
                className={cn(
                  errors.requiredSkills &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
                placeholder="SQL slamming"
              />
            </fieldset>
            <fieldset className="flex w-full flex-col gap-2">
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
                className={cn(
                  errors.benefits &&
                    "ring-2 ring-destructive ring-offset-2 focus-visible:ring-destructive"
                )}
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
          </CardContent>
          <CardFooter />
        </Card>
      </div>
      <div className="sticky top-0 hidden h-screen flex-1 flex-col items-end justify-between py-6 lg:flex">
        <Button
          disabled={isSubmitting}
          onClick={(...a) => void handleSubmit(onSubmit)(...a)}
        >
          {isSubmitting && (
            <Loader2Icon className="mr-2 h-4 w-4 animate-spin" />
          )}
          Save
        </Button>
        <ModeToggle />
      </div>
    </div>
  );
};

export { PostEditor };
