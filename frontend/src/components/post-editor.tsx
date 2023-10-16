"use client";

import * as React from "react";
import { useState } from "react";
import Link from "next/link";
import { TRANSFORMERS } from "@lexical/markdown";
import { AutoFocusPlugin } from "@lexical/react/LexicalAutoFocusPlugin";
import { LexicalComposer } from "@lexical/react/LexicalComposer";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import LexicalErrorBoundary from "@lexical/react/LexicalErrorBoundary";
import { HistoryPlugin } from "@lexical/react/LexicalHistoryPlugin";
import { LinkPlugin } from "@lexical/react/LexicalLinkPlugin";
import { ListPlugin } from "@lexical/react/LexicalListPlugin";
import { MarkdownShortcutPlugin } from "@lexical/react/LexicalMarkdownShortcutPlugin";
import { OnChangePlugin } from "@lexical/react/LexicalOnChangePlugin";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import { ChevronLeftIcon } from "lucide-react";
import TextareaAutosize from "react-textarea-autosize";
import { type Post } from "@/types/base/post";
import { initialConfig } from "@/lib/lexical";
import { CodeHighlightPlugin } from "./lexical/code-highlight-plugin";
import { ListMaxIndentLevelPlugin } from "./lexical/list-max-index-level-plugin";
import { ToolbarPlugin } from "./lexical/toolbar-plugin";
import { ModeToggle } from "./mode-toggle";
import { PostEditorSaveDialog } from "./post-editor-save-dialog";
import { Button } from "./ui/button";

type PostEditorProps = {
  postId: string;
  post: Post;
  // accessToken:
};

const PostEditor = ({ postId, post }: PostEditorProps) => {
  const editorState = post.description;
  const [description, setDescription] = useState<string>("{}");
  const [topic, setTopic] = useState(post?.topic ?? "Untitled Post");

  return (
    <div className="relative flex min-h-screen items-start md:container">
      <div className="sticky top-0 hidden h-screen flex-1 flex-col items-start py-6 lg:flex">
        <Button variant="ghost" asChild>
          <Link href="/dashboard/posts">
            <ChevronLeftIcon className="mr-2 h-4 w-4" />
            Back
          </Link>
        </Button>
      </div>
      <LexicalComposer initialConfig={{ ...initialConfig, editorState }}>
        <div className="relative mx-auto flex w-full max-w-3xl flex-col items-start">
          <div className="pointer-events-none sticky top-0 z-10 flex w-full justify-center bg-background pt-4">
            <ToolbarPlugin className="pointer-events-auto rounded-full border bg-background px-4 shadow-sm" />
          </div>
          <TextareaAutosize
            className="flex h-[5.625rem] w-full max-w-none resize-none appearance-none items-end bg-transparent px-8 pb-0.5 pt-12 text-4xl font-extrabold text-[#111827] scrollbar-hide focus:outline-none focus-visible:outline-none dark:text-[#ffffff]"
            value={topic}
            spellCheck={false}
            onChange={(e) => void setTopic(e.target.value)}
            aria-label="title"
          />
          <div className="relative w-full">
            <RichTextPlugin
              contentEditable={
                <ContentEditable
                  className="prose prose-green relative min-h-[28rem] max-w-none flex-1 bg-background px-8 pb-20 pt-8 dark:prose-invert focus:outline-none"
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
        </div>
        <HistoryPlugin />
        <AutoFocusPlugin />
        <CodeHighlightPlugin />
        <ListPlugin />
        <LinkPlugin />
        <ListMaxIndentLevelPlugin maxDepth={1} />
        <MarkdownShortcutPlugin transformers={TRANSFORMERS} />
        <OnChangePlugin
          onChange={(editorState) => {
            setDescription(JSON.stringify(editorState.toJSON()));
          }}
        />
      </LexicalComposer>
      <div className="sticky top-0 hidden h-screen flex-1 flex-col items-end justify-between py-6 lg:flex">
        <PostEditorSaveDialog
          postId={postId}
          post={post}
          topic={topic}
          description={description}
        />
        <ModeToggle />
      </div>
    </div>
  );
};

export { PostEditor };
