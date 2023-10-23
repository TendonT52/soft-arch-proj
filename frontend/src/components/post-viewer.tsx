"use client";

import * as React from "react";
import { useEffect, useState } from "react";
import { LexicalComposer } from "@lexical/react/LexicalComposer";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import LexicalErrorBoundary from "@lexical/react/LexicalErrorBoundary";
import { LinkPlugin } from "@lexical/react/LexicalLinkPlugin";
import { ListPlugin } from "@lexical/react/LexicalListPlugin";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import { format } from "date-fns";
import { CalendarIcon } from "lucide-react";
import { type Post } from "@/types/base/post";
import { editorConfig } from "@/lib/lexical";
import { formatDate, parsePeriod } from "@/lib/utils";
import { CodeHighlightPlugin } from "./lexical/code-highlight-plugin";
import { ListMaxIndentLevelPlugin } from "./lexical/list-max-index-level-plugin";
import { Loading } from "./loading";
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

type PostViewerProps = {
  post: Post & {
    owner: {
      id: string;
      name: string;
    };
    updatedAt: string;
    postId: string;
  };
};

const PostViewer = ({ post }: PostViewerProps) => {
  const {
    topic,
    description,
    openPositions,
    requiredSkills,
    benefits,
    howTo,
    owner,
    updatedAt,
  } = post;
  const period = parsePeriod(post.period);
  const [loading, setLoading] = useState(true);

  useEffect(() => void setLoading(false), []);

  return (
    <div className="relative mx-auto min-h-screen w-full max-w-3xl">
      <LexicalComposer
        initialConfig={{
          ...editorConfig,
          editable: false,
          editorState: description,
        }}
      >
        <div className="relative flex flex-col items-start">
          <div className="flex gap-6 px-8 pt-6">
            <div className="h-14 w-14 rounded-full bg-muted"></div>
            <div className="flex flex-col justify-end gap-2">
              <p className="text-muted-foreground">{owner.name}</p>
              <p className="text-sm leading-none text-muted-foreground">
                {formatDate(parseInt(updatedAt) * 1000)}
              </p>
            </div>
          </div>
          <div className="w-full max-w-none bg-transparent px-8 pb-0.5 pt-12 text-4xl font-extrabold text-[#111827] scrollbar-hide focus:outline-none focus-visible:outline-none dark:text-[#ffffff]">
            {topic}
          </div>
          {loading ? (
            <div className="relative flex min-h-[32rem] w-full items-center justify-center">
              <Loading />
            </div>
          ) : (
            <div className="relative w-full">
              <RichTextPlugin
                contentEditable={
                  <ContentEditable
                    className="prose prose-green relative min-h-[32rem] max-w-none flex-1 bg-background p-8 dark:prose-invert focus:outline-none"
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
        <CodeHighlightPlugin />
        <ListPlugin />
        <LinkPlugin />
        <ListMaxIndentLevelPlugin maxDepth={1} />
      </LexicalComposer>
      <Card className="rounded-none border-transparent border-t-border bg-background">
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
              <Input id="topic" value={topic} readOnly />
            </div>
            <div className="flex flex-1 flex-col gap-2">
              <Label
                htmlFor="period"
                className="w-full text-sm font-medium leading-none"
              >
                Period
              </Label>
              <Button
                variant={"outline"}
                className="justify-start text-left font-normal"
              >
                <CalendarIcon className="mr-2 h-4 w-4" />
                <span className="block">
                  {period?.from
                    ? period.to
                      ? `${format(period.from, "LLL dd, y")} - 
                  ${format(period.to, "LLL dd, y")}`
                      : format(period.from, "LLL dd, y")
                    : "Date range"}
                </span>
              </Button>
            </div>
          </fieldset>
          <fieldset className="flex w-full flex-col gap-2">
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
              id="openPositions"
              value={openPositions.join(" ")}
              placeholder="Top of the world"
              readOnly
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
              id="requiredSkills"
              value={requiredSkills.join(" ")}
              placeholder="SQL slamming"
              readOnly
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
              id="benefits"
              value={benefits.join(" ")}
              placeholder="Coffee"
              readOnly
            />
          </fieldset>
          <fieldset className="flex flex-col gap-2">
            <Label className="text-sm font-medium leading-none" htmlFor="howTo">
              How to apply
            </Label>
            <Textarea
              id="howTo"
              className="resize-none"
              value={howTo}
              placeholder="Run to the office like the flash ⚡"
              readOnly
            />
          </fieldset>
        </CardContent>
        <CardFooter />
      </Card>
    </div>
  );
};

export { PostViewer };
