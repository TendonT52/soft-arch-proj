"use client";

import * as React from "react";
import { useEffect, useState } from "react";
import { LexicalComposer } from "@lexical/react/LexicalComposer";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import LexicalErrorBoundary from "@lexical/react/LexicalErrorBoundary";
import { LinkPlugin } from "@lexical/react/LexicalLinkPlugin";
import { ListPlugin } from "@lexical/react/LexicalListPlugin";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import { type Post } from "@/types/base/post";
import { initialConfig } from "@/lib/lexical";
import { formatDate } from "@/lib/utils";
import { CodeHighlightPlugin } from "./lexical/code-highlight-plugin";
import { ListMaxIndentLevelPlugin } from "./lexical/list-max-index-level-plugin";
import { Loading } from "./loading";

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
  const { topic, description, owner, updatedAt } = post;
  const [loading, setLoading] = useState(true);

  useEffect(() => void setLoading(false), []);

  return (
    <LexicalComposer
      initialConfig={{
        ...initialConfig,
        editable: false,
        editorState: description,
      }}
    >
      <div className="relative mx-auto flex w-full max-w-3xl flex-col items-start">
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
          <div className="relative flex min-h-[28rem] w-full items-center justify-center">
            <Loading />
          </div>
        ) : (
          <div className="relative w-full">
            <RichTextPlugin
              contentEditable={
                <ContentEditable
                  className="prose prose-green relative min-h-[28rem] max-w-none flex-1 bg-background p-8 dark:prose-invert focus:outline-none"
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
  );
};

export { PostViewer };
