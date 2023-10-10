"use client";

import * as React from "react";
import { LexicalComposer } from "@lexical/react/LexicalComposer";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import LexicalErrorBoundary from "@lexical/react/LexicalErrorBoundary";
import { LinkPlugin } from "@lexical/react/LexicalLinkPlugin";
import { ListPlugin } from "@lexical/react/LexicalListPlugin";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import { initialConfig } from "@/lib/lexical";
import { CodeHighlightPlugin } from "./lexical/code-highlight-plugin";
import { ListMaxIndentLevelPlugin } from "./lexical/list-max-index-level-plugin";

type Post = {
  topic?: string;
  description?: string;
};

type PostEditorProps = {
  post?: Post;
};

const PostViewer = ({
  post: { topic: title, description: description } = {},
}: PostEditorProps) => {
  return (
    <LexicalComposer
      initialConfig={{
        ...initialConfig,
        editable: false,
        editorState: description,
      }}
    >
      <div className="relative mx-auto flex w-full max-w-3xl flex-col items-start">
        <div className="w-full max-w-none bg-transparent px-8 pb-0.5 pt-12 text-4xl font-extrabold text-[#111827] scrollbar-hide focus:outline-none focus-visible:outline-none dark:text-[#ffffff]">
          {title}
        </div>
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
      </div>
      <CodeHighlightPlugin />
      <ListPlugin />
      <LinkPlugin />
      <ListMaxIndentLevelPlugin maxDepth={1} />
    </LexicalComposer>
  );
};
PostViewer.displayName = "PostViewer";

export { PostViewer };
